// breweryReviews/beer_handler/breweryReviews_handler.go
package breweryReviews_handler

import (
	"Beer.app/breweryReviews/repository"
	"Beer.app/breweryReviews/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BreweryReviewHandler struct {
	service service.BreweryReviewService
}

func NewReviewHandler(service service.BreweryReviewService) *BreweryReviewHandler {
	return &BreweryReviewHandler{service: service}
}

func BreweryReviewsRouter(e *echo.Echo, db *gorm.DB) {
	reviewRepo := repository.NewBreweryReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := NewReviewHandler(reviewService)

	e.GET("/brewery/reviews/:id", reviewHandler.GetBreweryReview)
	e.POST("/brewery/reviews", reviewHandler.CreateBreweryReview)
}

func (h *BreweryReviewHandler) GetBreweryReview(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid review ID"})
	}

	review, err := h.service.GetBreweryReview(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Review not found"})
	}
	return c.JSON(http.StatusOK, review)
}

func (h *BreweryReviewHandler) CreateBreweryReview(c echo.Context) error {
	review := new(models.BreweryReview)
	if err := c.Bind(review); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	if err := h.service.CreateBreweryReview(review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create review"})
	}
	return c.JSON(http.StatusCreated, review)
}
