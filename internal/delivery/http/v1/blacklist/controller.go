package blacklist

import (
	"github.com/labstack/echo/v4"

	"github.com/MaxFando/anti-bruteforce/internal/domain/network"
	"github.com/MaxFando/anti-bruteforce/internal/usecase/blacklist"
)

type Controller struct {
	uc *blacklist.UseCase
}

func NewController(usecase *blacklist.UseCase) *Controller {
	return &Controller{
		uc: usecase,
	}
}

type addIPRequest struct {
	IP   string `json:"ip" query:"ip" validate:"required"`
	Mask string `json:"mask" query:"mask" validate:"required"`
}

func (ctr *Controller) AddIP(c echo.Context) error {
	ctx := c.Request().Context()

	request := new(addIPRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(500, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(422, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	payload, err := network.NewIPNetwork(request.IP, request.Mask)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	err = ctr.uc.AddIP(ctx, payload)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	return c.JSON(200, map[string]interface{}{"ok": true})
}

type removeIPRequest struct {
	IP   string `json:"ip" query:"ip" validate:"required"`
	Mask string `json:"mask" query:"mask" validate:"required"`
}

func (ctr *Controller) RemoveIP(c echo.Context) error {
	ctx := c.Request().Context()

	request := new(removeIPRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(500, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(422, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	payload, err := network.NewIPNetwork(request.IP, request.Mask)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	err = ctr.uc.RemoveIP(ctx, payload)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	return c.JSON(200, map[string]interface{}{"ok": true})
}

func (ctr *Controller) GetIPList(c echo.Context) error {
	ctx := c.Request().Context()

	ipList, err := ctr.uc.GetIPList(ctx)
	if err != nil {
		return c.JSON(500, map[string]interface{}{"ok": false, "error": err.Error()})
	}

	return c.JSON(200, map[string]interface{}{"ok": true, "data": ipList})
}
