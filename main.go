// main.go
package main

import (
	"Beer.app/brewery/brewery_handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Register user routes
	brewery_handler.UserRouter(e)
	brewery_handler.ProductRouter(e)
	brewery_handler.ReviewRouter(e)
	brewery_handler.BreweryRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
