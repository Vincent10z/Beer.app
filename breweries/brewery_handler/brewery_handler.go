// users/account_handler/breweries_handler.go
package brewery_handler

import (
	"Beer.app/breweries/repository"
	"Beer.app/breweries/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type BreweryHandler struct {
	service service.BreweryService
}

func NewBreweryHandler(service service.BreweryService) *BreweryHandler {
	return &BreweryHandler{service: service}
}

func BreweryRouter(e *echo.Echo, db *gorm.DB) {
	breweryRepo := repository.NewBreweryRepository(db)
	breweryService := service.NewBreweryService(breweryRepo)
	breweryHandler := NewBreweryHandler(breweryService)

	e.GET("/breweries/:id", breweryHandler.GetBreweryByID)
	e.POST("/breweries", breweryHandler.CreateBrewery)
}

func (h *BreweryHandler) GetBreweryByID(c echo.Context) error {
	id := c.Param("id")

	brewery, err := h.service.GetBrewery(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Brewery not found"})
	}

	return c.JSON(http.StatusOK, brewery)
}

func (h *BreweryHandler) CreateBrewery(c echo.Context) error {
	brewery := &models.Brewery{}
	if err := c.Bind(brewery); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	if err := h.service.CreateBrewery(brewery); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create breweries"})
	}

	return c.JSON(http.StatusCreated, brewery)
}
