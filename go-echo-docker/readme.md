# [Golang] Dockerfile

---

### Create ``main.go``
```sh
$ mkdir project
$ touch main.go
```

### Init Go Module

```sh
$ go mod init
```

### Import

```sh
package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)
```

### Function main

```sh
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3000"))
}
```

### Dockerfile

```sh
FROM golang:1.14.0
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
```

### Docker build and run

```sh
$ docker build -t go-echo-docker .
$ docker run -dp 3000:3000 --name go-echo go-echo-docker
```

เปิด browser แล้วเข้า http://localhost:3000

---


![Screen Shot 2563-11-02 at 11.08.23.png](./image.png)