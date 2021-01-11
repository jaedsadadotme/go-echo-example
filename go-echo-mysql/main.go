package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	env := godotenv.Load() //Load .env file
	if env != nil {
		fmt.Print(env)
	}

	h := UserHandler{}
	h.Initialize()

	e.GET("/", h.GetAll)
	e.GET("/:id", h.GetOne)
	e.POST("/", h.Create)
	e.PUT("/:id", h.Update)
	e.DELETE("/:id", h.Destroy)
	e.Logger.Fatal(e.Start(":1323"))
}
