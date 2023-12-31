package main

import (
	"context"
	"fmt"
	"github.com/MaxFando/anti-bruteforce/internal/app/cli"
	"github.com/MaxFando/anti-bruteforce/internal/app/grpcapi"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/MaxFando/anti-bruteforce/internal/app/http"
	"github.com/MaxFando/anti-bruteforce/internal/config"
	"github.com/MaxFando/anti-bruteforce/internal/providers"
	"github.com/MaxFando/anti-bruteforce/internal/store/postgres"
	"github.com/MaxFando/anti-bruteforce/pkg/utils"
)

func main() {
	viper.AutomaticEnv()

	utils.InitializeLogger()

	ctx, cancel := context.WithCancel(context.Background())
	config.InitializeConfig()

	postgresDB := postgres.New()
	err := postgresDB.Connect(ctx, "default", config.Config.Database)
	if err != nil {
		panic(err)
	}

	repositoryProvider := providers.NewRepositoryProvider(postgresDB)
	repositoryProvider.BootPrefixProviderContract()

	serviceProvider := providers.NewServiceProvider()
	serviceProvider.RegisterDependencies(repositoryProvider)

	useCaseProvider := providers.NewUseCaseProvider()
	useCaseProvider.RegisterDependencies(serviceProvider)

	ctx = context.WithValue(ctx, providers.ServiceProviderKey, serviceProvider)
	ctx = context.WithValue(ctx, providers.UseCaseProviderKey, useCaseProvider)

	httpServer := http.NewHTTPServer(http.NewHandler(ctx), ":"+config.Config.Listen.Port)
	httpServer.Serve()
	utils.Logger.Info("Приложение стартовало в режиме", zap.String("log_level", config.Config.AppConfig.LogLevel))
	utils.Logger.Info("На порту " + config.Config.Listen.Port)

	grpcServer := grpcapi.NewServer(ctx)
	grpcServer.Serve()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	cmd := cli.NewCmd()
	cmd.Run(ctx, interrupt)

	select {
	case s := <-interrupt:
		utils.Logger.Info("app - Run - signal: " + s.String())
	case errHTTP := <-httpServer.Notify():
		utils.Logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", errHTTP))
	case errCmd := <-cmd.Notify():
		utils.Logger.Error(fmt.Errorf("app - Run - cmd.Notify: %w", errCmd))
	case errGrpc := <-grpcServer.Notify():
		utils.Logger.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", errGrpc))
	}

	cancel()
	_ = httpServer.Shutdown()
	grpcServer.Shutdown()

	utils.Logger.Info("Приложение завершило работу")
}
