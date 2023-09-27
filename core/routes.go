package core

import (
	"restapi/handlers"
	"restapi/middlewares"
)

func (api *Api) Routes() {
	r := api.App.Group("/api/v1")

	r.Use(middlewares.Handler(api.Db, api.Log))
	r.Use(middlewares.ErrorHandler())
	{
		r.GET("/ping", handlers.Ping)
	}
}
