package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yourusername/yourproject/handlers"
)

func main() {
	e := echo.New()

	e.GET("/doctors", handlers.GetDoctors)
	e.GET("/doctors/:id", handlers.GetDoctorByID)
	e.POST("/doctors", handlers.CreateDoctor)
	e.PUT("/doctors/:id", handlers.UpdateDoctor)
	e.DELETE("/doctors/:id", handlers.DeleteDoctor)

	e.Logger.Fatal(e.Start(":8080"))
}
