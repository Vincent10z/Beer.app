// main.go
package main

import (
	"Beer.app/users/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Register user routes
	http.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
