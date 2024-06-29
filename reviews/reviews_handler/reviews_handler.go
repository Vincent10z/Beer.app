// reviews/product_handler/reviews_handler.go
package reviews_handler

import (
	"Beer.app/models"
	"Beer.app/reviews/repository"
	"Beer.app/reviews/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

func ReviewRouter(e *echo.Echo) {
	reviewRepo := repository.NewReviewRepository()
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := NewReviewHandler(reviewService)

	e.GET("/reviews/:id", reviewHandler.GetReview)
	e.POST("/reviews", reviewHandler.CreateReview)
}

func (h *ReviewHandler) GetReview(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid review ID"})
	}

	review, err := h.service.GetReview(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Review not found"})
	}
	return c.JSON(http.StatusOK, review)
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	review := new(models.Review)
	if err := c.Bind(review); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	if err := h.service.CreateReview(review); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create review"})
	}
	return c.JSON(http.StatusCreated, review)
}
