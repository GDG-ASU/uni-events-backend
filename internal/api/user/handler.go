package user

import (
	"net/http"
	"uni-events-backend/internal/models"
	"uni-events-backend/pkg/utils"

	"github.com/labstack/echo/v4"

	"uni-events-backend/internal/service"
)

type Handler struct {
	service service.UserService
}

func NewHandler(userService service.UserService) *Handler {
	return &Handler{
		service: userService,
	}
}

func (h *Handler) GetMe(c echo.Context) error {
	clerkID := utils.GetClerkUserID(c)

	user, err := h.service.GetUserByClerkID(c.Request().Context(),clerkID)
	if err != nil {
		newUser := &models.User{
			ClerkID: clerkID,
			Email:   c.Get("email").(string),
			Role:    "student",
		}
		createdUser, err := h.service.CreateUserIfNotExists(c.Request().Context(), newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}
		
		return c.JSON(http.StatusOK, createdUser)
	}

	return c.JSON(http.StatusOK, user)
}
