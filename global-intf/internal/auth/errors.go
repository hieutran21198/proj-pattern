package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidCredential = echo.NewHTTPError(http.StatusBadRequest, "user is not found")
)
