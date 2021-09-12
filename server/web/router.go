package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

// Option to allow config an echo instance
type Option func(*echo.Echo)

// NewRouter creates a new echo instance, uses provided options and sets the default middleware
func NewRouter(options ...Option) *echo.Echo {
	e := echo.New()

	// use provided options
	for _, opt := range options {
		opt(e)
	}

	// append default middleware
	e.Use(middleware.Recover())
	return e
}