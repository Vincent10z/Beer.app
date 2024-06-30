// users/beer_handler/breweryReviews_handler.go
package users_handler

import (
	"Beer.app/models"
	"Beer.app/users/repository"
	"Beer.app/users/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(UserService service.UserService) *UserHandler {
	return &UserHandler{UserService: UserService}
}

// UserRouter to register user routes
func UserRouter(e *echo.Echo, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
}

// GetUser handles `GET /users/:id`
func (h *UserHandler) GetUser(c echo.Context) error {
	user, err := h.UserService.GetUser(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, user)
}

// CreateUser handles `POST /users`
func (h *UserHandler) CreateUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.UserService.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}
