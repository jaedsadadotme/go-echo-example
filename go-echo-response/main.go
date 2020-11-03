package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// ResponseJSON ...
func ResponseJSON(c echo.Context) error {
	dataReponse := map[string]string{
		"message": "Hello",
	}
	return c.JSON(http.StatusOK, dataReponse)
}

// ResponseString ...
func ResponseString(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

// ResponseHTML ...
func ResponseHTML(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello</h1>")
}

// ResponseStream ...
func ResponseStream(c echo.Context) error {
	file, err := os.Open("dog.png")
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, "image/png", file)
}

func main() {
	e := echo.New()
	e.GET("/json", ResponseJSON)
	e.GET("/string", ResponseString)
	e.GET("/html", ResponseHTML)
	e.GET("/file", ResponseStream)
	e.Logger.Fatal(e.Start(":3000"))
}
