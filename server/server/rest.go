package server

import (
	"github.com/labstack/echo/v4"
)

const apiPrefix = "/api/v1"

type Rest struct {
	Router  *echo.Echo
	Handler *handler.Handler
}

func (r *Rest) Route() {
	route := r.Router.Group(apiPrefix)
	
	route.GET("/receiver/get_settings", r.GetReceiverSettings)
	

	route.POST("/receiver/set_settings",r.SetSettings)
	route.POST("/receiver/set_mode",r.SetMode)
	route.POST("/receiver/set_freq",r.SetFreq)
	}