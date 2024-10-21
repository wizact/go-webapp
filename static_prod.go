//go:build !dev

package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed assets/**
var PublicFS embed.FS

func publicMiddleware() echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "assets",
		Browse:     false,
		Filesystem: http.FS(PublicFS),
	})
}
