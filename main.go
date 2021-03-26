package main

import (
	"net/http"

	"github.com/ANijat/creature-crud/controller"
	"github.com/ANijat/creature-crud/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", hello)
	e.GET("/creature/", controller.GetCreatures)
	e.GET("/creature/:id", controller.GetCreaturesByid)
	e.PUT("/creature/:id", controller.UpdateCreature)
	e.POST("/creature/", controller.CreateCreature)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}
