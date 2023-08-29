package main

import (
	"github.com/Takeasoul/Task-manager/api"
	"github.com/Takeasoul/Task-manager/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/static", "web/static")

	// Load templates
	e.Renderer = web.NewTemplateRenderer()

	// Initialize API routes
	api.InitRoutes(e)

	// Initialize web routes
	web.InitRoutes(e)

	// Start server
	e.Start(":8080")
}
