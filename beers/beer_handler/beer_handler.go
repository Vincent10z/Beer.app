package beer_handler

import (
	"Beer.app/beers/repository"
	"Beer.app/beers/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BeerHandler struct {
	service service.BeerService
}

func BeerRouter(e *echo.Echo) {
	beerRepo := repository.NewBeerRepository()
	beerService := service.NewBeerService(beerRepo)
	beerHandler := NewBeerHandler(beerService)

	e.GET("/beers/:id", beerHandler.GetBeer)
	e.POST("/beers", beerHandler.CreateBeer)
}

func NewBeerHandler(service service.BeerService) *BeerHandler {
	return &BeerHandler{service: service}
}

func (h *BeerHandler) GetBeer(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid beer ID"})
	}

	beer, err := h.service.GetBeer(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Beer not found"})
	}
	return c.JSON(http.StatusOK, beer)
}

func (h *BeerHandler) CreateBeer(c echo.Context) error {
	beer := models.Beer{}
	if err := c.Bind(&beer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	if err := h.service.CreateBeer(&beer); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create beer"})
	}
	return c.JSON(http.StatusCreated, beer)
}
