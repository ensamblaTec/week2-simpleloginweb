package handler

import (
	"fmt"
	"net/http"

	"github.com/ensamblaTec/learning/week/problema3/pkg/models"
	"github.com/labstack/echo/v4"
)

func LoginTemplate(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("pass")
	verify := models.CreateLogin("admin@example.com", "1234")
	login := models.CreateLogin(email, password)
	if email == verify.Email && login.Password == verify.Password {
		login.Logged = true
	}
	fmt.Println(login.Email, login.Logged, login.Password)
	fmt.Println(verify.Email, verify.Logged, verify.Password)
	return c.Render(http.StatusOK, "completeMsg.html", login)
}
