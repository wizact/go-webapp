//go:build dev

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func publicMiddleware() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "assets",
		Browse: true,
	})
}
