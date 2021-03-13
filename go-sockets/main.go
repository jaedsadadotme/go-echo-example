package main

import (
	"fmt"
	"time"

	// "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

// func gorilla(c echo.Context) error {
// 	var (
// 		upgrader = websocket.Upgrader{}
// 	)

// 	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer ws.Close()

// 	for {
// 		// Read
// 		_, msg, err := ws.ReadMessage()
// 		if err != nil {
// 			c.Logger().Error(err)
// 		}
// 		fmt.Printf("%s\n", msg)

// 		// Write
// 		if string(msg) == "time" {
// 			error := ws.WriteMessage(websocket.TextMessage, []byte(time.Now().String()))
// 			if error != nil {
// 				c.Logger().Error(err)
// 			}
// 		} else {
// 			error := ws.WriteMessage(websocket.TextMessage, msg)
// 			if error != nil {
// 				c.Logger().Error(err)
// 			}
// 		}

// 	}
// }

func useNet(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)

			// Write
			if string(msg) == "time" {
				error := websocket.Message.Send(ws, time.Now().String())
				if error != nil {
					c.Logger().Error(error)
				}
			} else {
				error := websocket.Message.Send(ws, msg)
				if error != nil {
					c.Logger().Error(error)
				}
			}

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public")
	e.GET("/ws", useNet)
	e.Logger.Fatal(e.Start(":1323"))
}
