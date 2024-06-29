package beer_handler

import (
	"Beer.app/beers/repository"
	"Beer.app/beers/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BeerHandler struct {
	service service.BeerService
}

func BeerRouter(e *echo.Echo, db *gorm.DB) {
	beerRepo := repository.NewBeerRepository(db)
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
		return c.JSON(http.StatusBadRequest, err)
	}

	beer, err := h.service.GetBeer(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if beer == nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, beer)
}

func (h *BeerHandler) CreateBeer(c echo.Context) error {
	beer := models.Beer{}
	if err := c.Bind(&beer); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.service.CreateBeer(&beer); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, beer)
}
