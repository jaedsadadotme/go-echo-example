package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)

		// Write
		if string(msg) == "time" {
			error := ws.WriteMessage(websocket.TextMessage, []byte(time.Now().String()))
			if error != nil {
				c.Logger().Error(err)
			}
		} else {
			error := ws.WriteMessage(websocket.TextMessage, msg)
			if error != nil {
				c.Logger().Error(err)
			}
		}

	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
