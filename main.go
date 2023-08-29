package main

import (
	"github.com/Takeasoul/Task-manager/web"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files from "public" directory
	e.Static("/static", "public")

	// Load templates
	e.Renderer = web.NewTemplateRenderer()

	// Initialize web routes
	web.InitRoutes(e)

	// Start server
	e.Start(":8080")
}
