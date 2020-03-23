package router

import (
	"spiros/server/handlers"
	"spiros/server/helper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ClientGroup group of client's endpoints
func ClientGroup(e *echo.Echo) {
	g := e.Group("/client")
	g.Use(middleware.BasicAuth(helper.ValidateClient))

	g.POST("/login", handlers.Login)
}
