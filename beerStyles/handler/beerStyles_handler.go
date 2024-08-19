// beerStyles/handler/beer_style_handler.go
package beer_style_handler

import (
	"Beer.app/beerStyles/repository"
	"Beer.app/beerStyles/service"
	"Beer.app/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type BeerStyleHandler struct {
	service service.BeerStyleService
}

func NewBeerStyleHandler(service service.BeerStyleService) *BeerStyleHandler {
	return &BeerStyleHandler{service: service}
}

func BeerStyleRouter(e *echo.Echo, db *gorm.DB) {
	beerStyleRepo := repository.NewBeerStyleRepository(db)
	beerStyleService := service.NewBeerStyleService(beerStyleRepo)
	beerStyleHandler := NewBeerStyleHandler(beerStyleService)

	e.GET("/beer-styles/:id", beerStyleHandler.GetBeerStyle)
	e.POST("/beer-styles", beerStyleHandler.CreateBeerStyle)
	e.PUT("/beer-styles/:id", beerStyleHandler.UpdateBeerStyle)
	e.DELETE("/beer-styles/:id", beerStyleHandler.DeleteBeerStyle)
	e.GET("/beer-styles", beerStyleHandler.ListBeerStyles)
	e.POST("/beer-styles/bulk", beerStyleHandler.CreateMultipleBeerStyles)
}

func (h *BeerStyleHandler) GetBeerStyle(c echo.Context) error {
	id := c.Param("id")
	style, err := h.service.GetBeerStyle(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Beer style not found"})
	}
	return c.JSON(http.StatusOK, style)
}

func (h *BeerStyleHandler) CreateBeerStyle(c echo.Context) error {
	style := new(models.BeerStyle)
	if err := c.Bind(style); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := h.service.CreateBeerStyle(style); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create beer style"})
	}
	return c.JSON(http.StatusCreated, style)
}

func (h *BeerStyleHandler) UpdateBeerStyle(c echo.Context) error {
	id := c.Param("id")
	style := new(models.BeerStyle)
	if err := c.Bind(style); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	style.ID = id
	if err := h.service.UpdateBeerStyle(style); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update beer style"})
	}
	return c.JSON(http.StatusOK, style)
}

func (h *BeerStyleHandler) DeleteBeerStyle(c echo.Context) error {
	id := c.Param("id")
	if err := h.service.DeleteBeerStyle(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete beer style"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *BeerStyleHandler) ListBeerStyles(c echo.Context) error {
	styles, err := h.service.ListBeerStyles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to list beer styles"})
	}
	return c.JSON(http.StatusOK, styles)
}

func (h *BeerStyleHandler) CreateMultipleBeerStyles(c echo.Context) error {
	var styles []*models.BeerStyle
	if err := c.Bind(&styles); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := h.service.CreateMultipleBeerStyles(styles); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create beer styles"})
	}
	return c.JSON(http.StatusCreated, styles)
}
