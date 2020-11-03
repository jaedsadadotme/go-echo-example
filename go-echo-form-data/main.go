package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name" form:"name"`
}

func addName(c echo.Context) (err error) {
	u := new(User)
	if err = c.Bind(u); err != nil {
		return
	}
	return c.JSON(http.StatusOK, u)
}

func main() {
	e := echo.New()
	e.POST("/adduser", addName)
	e.Logger.Fatal(e.Start(":3000"))
}
