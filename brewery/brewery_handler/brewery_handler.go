// users/product_handler/breweries_handler.go
package brewery_handler

import (
	"Beer.app/brewery/repository"
	"Beer.app/brewery/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BreweryHandler struct {
	service service.BreweryService
}

func NewBreweryHandler(service service.BreweryService) *BreweryHandler {
	return &BreweryHandler{service: service}
}

func BreweryRouter(e *echo.Echo) {
	breweryRepo := repository.NewBreweryRepository()
	breweryService := service.NewBreweryService(breweryRepo)
	breweryHandler := NewBreweryHandler(breweryService)

	e.GET("/brewery/:id", breweryHandler.GetBreweryByID)
	e.POST("/brewery", breweryHandler.CreateBrewery)
}

func (h *BreweryHandler) GetBreweryByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid brewery ID"})
	}

	brewery, err := h.service.GetBrewery(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Brewery not found"})
	}

	return c.JSON(http.StatusOK, brewery)
}

func (h *BreweryHandler) CreateBrewery(c echo.Context) error {
	brewery := new(models.Brewery)
	if err := c.Bind(brewery); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	if err := h.service.CreateBrewery(brewery); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create brewery"})
	}

	return c.JSON(http.StatusCreated, brewery)
}
