package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"global-intf/internal/dto"
	"global-intf/internal/intf"
)

type HTTP struct {
	srv intf.AuthService
}

// TODO: NewHTTP ->

func (h *HTTP) Login(c echo.Context) error {
	var (
		ctx = c.Request().Context()
		req = &dto.LoginRequest{}
		res *dto.LoginResponse
		err error
	)

	// TODO: test binding rule by validator (attached to echo server)
	if err = c.Bind(req); err != nil {
		return err
	}

	res, err = h.srv.Login(ctx, req)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
