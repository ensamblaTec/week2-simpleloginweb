package main

import (
	"github.com/ensamblaTec/learning/week/problema3/pkg/handler"
	"github.com/ensamblaTec/learning/week/problema3/pkg/models"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true

	template := models.LoadTemplates()

	e.Renderer = template
	// handlers
	e.GET("/", handler.LoginTemplate)

	e.POST("/auth/login", handler.Login)

	e.Logger.Fatal(e.Start(":80"))
}
