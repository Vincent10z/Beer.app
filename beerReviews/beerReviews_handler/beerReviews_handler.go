// breweryReviews/beer_handler/breweryReviews_handler.go
package beerReviews_handler

import (
	"Beer.app/beerReviews/repository"
	"Beer.app/beerReviews/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BeerReviewHandler struct {
	service service.BeerReviewService
}

func NewBeerReviewHandler(service service.BeerReviewService) *BeerReviewHandler {
	return &BeerReviewHandler{service: service}
}

func BeerReviewRouter(e *echo.Echo, db *gorm.DB) {
	beerReviewRepo := repository.NewBeerReviewRepository(db)
	beerReviewService := service.NewBeerReviewService(beerReviewRepo)
	beerReviewHandler := NewBeerReviewHandler(beerReviewService)

	e.GET("/beerReviews/:id", beerReviewHandler.GetBeerReview)
	e.POST("/beerReviews", beerReviewHandler.CreateBeerReview)
}

func (h *BeerReviewHandler) GetBeerReview(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid review ID"})
	}

	review, err := h.service.GetBeerReview(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Review not found"})
	}
	return c.JSON(http.StatusOK, review)
}

func (h *BeerReviewHandler) CreateBeerReview(c echo.Context) error {
	review := new(models.BeerReview)
	if err := c.Bind(review); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	if err := h.service.CreateBeerReview(review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create review"})
	}
	return c.JSON(http.StatusCreated, review)
}
