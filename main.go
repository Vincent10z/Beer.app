// main.go
package main

import (
	"Beer.app/account/account_handler"
	"Beer.app/beerReviews/beerReviews_handler"
	"Beer.app/beerStyles/handler"
	"Beer.app/beers/beer_handler"
	"Beer.app/breweries/brewery_handler"
	"Beer.app/breweryReviews/breweryReviews_handler"
	"Beer.app/database"
	"Beer.app/users/users_handler"
	"github.com/labstack/echo/v4"
)

func main() {
	db := database.ConnectDB()

	e := echo.New()

	// Register routes
	account_handler.AccountRouter(e, db)
	users_handler.UserRouter(e, db)
	beer_handler.BeerRouter(e, db)
	beerReviews_handler.BeerReviewRouter(e, db)
	breweryReviews_handler.BreweryReviewsRouter(e, db)
	brewery_handler.BreweryRouter(e, db)
	beer_style_handler.BeerStyleRouter(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
