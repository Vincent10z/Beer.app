package beer_handler

import (
	"Beer.app/beers/repository"
	"Beer.app/beers/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type BeerHandler struct {
	service service.BeerService
}

func NewBeerHandler(service service.BeerService) *BeerHandler {
	return &BeerHandler{service: service}
}

func BeerRouter(e *echo.Echo, db *gorm.DB) {
	beerRepo := repository.NewBeerRepository(db)
	beerService := service.NewBeerService(beerRepo)
	beerHandler := NewBeerHandler(beerService)

	e.POST("/beers", beerHandler.CreateBeer)
	e.GET("/beers/:id", beerHandler.GetBeer)
	e.PUT("/beers/:id", beerHandler.UpdateBeer)
	e.DELETE("/beers/:id", beerHandler.DeleteBeer)
	e.GET("/beers", beerHandler.ListBeers)
}

func (h *BeerHandler) CreateBeer(c echo.Context) error {
	beer := new(models.Beer)
	if err := c.Bind(beer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	beer, err := h.service.CreateBeer(beer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create beer"})
	}

	return c.JSON(http.StatusCreated, beer)
}

func (h *BeerHandler) GetBeer(c echo.Context) error {
	id := c.Param("id")
	beer, err := h.service.GetBeer(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Beer not found"})
	}
	return c.JSON(http.StatusOK, beer)
}

func (h *BeerHandler) UpdateBeer(c echo.Context) error {
	id := c.Param("id")
	beer := new(models.Beer)
	if err := c.Bind(beer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	beer.ID = id

	updatedBeer, err := h.service.UpdateBeer(beer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update beer"})
	}

	return c.JSON(http.StatusOK, updatedBeer)
}

func (h *BeerHandler) DeleteBeer(c echo.Context) error {
	id := c.Param("id")
	err := h.service.DeleteBeer(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete beer"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *BeerHandler) ListBeers(c echo.Context) error {
	beers, err := h.service.ListBeers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to list beers"})
	}
	return c.JSON(http.StatusOK, beers)
}
