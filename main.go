package main

import (
	"net/http"
	"time"

	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/wizact/go-webapp/views/pages/home"

	"github.com/wizact/go-webapp/views/pages/about"
	"github.com/wizact/go-webapp/views/pages/product"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		component := home.Index("John!!")
		return Render(c, http.StatusOK, component)
	})
	e.GET("/about", func(c echo.Context) error {
		component := about.About("John!!")
		return Render(c, http.StatusOK, component)
	})
	e.GET("/products", func(c echo.Context) error {
		component := product.List("John!!")
		return Render(c, http.StatusOK, component)
	})

	group := e.Group("assets")
	group.Use(publicMiddleware())

	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := e.Start(":1323")
		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started...")
	<-killSig

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown completed")
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
