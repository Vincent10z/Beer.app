package http

import (
	"Beer.app/models"
	"Beer.app/product/repository"
	"Beer.app/product/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service service.ProductService
}

func productRouter(e *echo.Echo) {
	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	e.GET("/users/:id", productHandler.GetProduct)
	e.POST("/users", productHandler.CreateProduct)
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid user ID"})
	}

	user, err := h.service.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := models.Product{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}

	if err := h.service.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}
	return c.JSON(http.StatusCreated, product)
}
