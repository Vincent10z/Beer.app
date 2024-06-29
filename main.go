// main.go
package main

import (
	"Beer.app/beerReviews/beerReviews_handler"
	"Beer.app/beers/beer_handler"
	"Beer.app/breweries/brewery_handler"
	"Beer.app/users/users_handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Register user routes
	users_handler.UserRouter(e)
	beerReviews_handler.BeerReviewRouter(e)
	brewery_handler.BreweryRouter(e)
	beer_handler.BeerRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
