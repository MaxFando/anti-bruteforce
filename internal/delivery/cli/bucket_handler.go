package cli

import (
	"context"
	"fmt"
	"github.com/MaxFando/anti-bruteforce/internal/domain/network"
)

func (c *CommandLineInterface) bucketHandler(ctx context.Context, setCommand []string) {
	if len(setCommand) != 4 {
		return
	}
	if setCommand[1] == "reset" {
		request, err := network.NewRequest(
			setCommand[2],
			"",
			setCommand[3],
		)
		if err != nil {
			fmt.Printf("bootExecutor - bucket reset: %s", err.Error())
			return
		}

		c.resetBucket(ctx, request)
	} else {
		fmt.Println("unknown command")
	}
}

func (c *CommandLineInterface) resetBucket(ctx context.Context, request network.Request) {
	isLoginReset, isIPReset, err := c.bucketUseCase.Reset(ctx, request.Login, request.IP.String())
	if err != nil {
		fmt.Printf("service error: %v \n", err)
		return
	}

	if !isLoginReset {
		fmt.Printf("login: %v has not been reseted\n", request.Login)
	} else {
		fmt.Printf("login: %v has been reseted\n", request.Login)
	}

	if !isIPReset {
		fmt.Printf("ip: %v has not been reseted\n", request.IP)
	} else {
		fmt.Printf("ip: %v has been reseted\n", request.IP)
	}
}
