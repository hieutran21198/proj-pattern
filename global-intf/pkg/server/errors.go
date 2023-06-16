package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInternalDatabaseError = echo.NewHTTPError(http.StatusInternalServerError, "internal database error")
)
