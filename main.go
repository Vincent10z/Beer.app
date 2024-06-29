// main.go
package main

import (
	"Beer.app/beerReviews/beerReviews_handler"
	"Beer.app/beers/beer_handler"
	"Beer.app/breweries/brewery_handler"
	"Beer.app/breweryReviews/breweryReviews_handler"
	"Beer.app/database"
	"Beer.app/users/users_handler"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	e := echo.New()

	// Register user routes
	users_handler.UserRouter(e)
	beer_handler.BeerRouter(e)
	beerReviews_handler.BeerReviewRouter(e)
	breweryReviews_handler.BreweryReviewsRouter(e)
	brewery_handler.BreweryRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
