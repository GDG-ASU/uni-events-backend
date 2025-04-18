package club

import (
	"net/http"
	"strconv"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/service"
	"uni-events-backend/pkg/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service service.ClubService
	userService service.UserService
}

func NewHandler(clubService service.ClubService, userService service.UserService) *Handler {
	return &Handler{
		service:     clubService,
		userService: userService,
	}
}

type CreateClubRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateClubRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}


func (h *Handler) CreateClub(c echo.Context) error {
	clerkID := utils.GetClerkUserID(c)

	// You should already have a method to get the UserID from ClerkID
	user, err := h.userService.GetUserByClerkID(c.Request().Context(), clerkID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "User not found"})
	}

	var req CreateClubRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	club := &models.Club{
		Name:        req.Name,
		Description: req.Description,
		Owners: []models.ClubOwner{
			{UserID: user.ID}, // Link user as the owner
		},
	}

	createdClub, err := h.service.CreateClub(c.Request().Context(), club)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdClub)
}

func (h *Handler) UpdateClub(c echo.Context) error {
	clerkID := utils.GetClerkUserID(c)
	user, err := h.userService.GetUserByClerkID(c.Request().Context(), clerkID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
	}

	clubIDStr := c.Param("id")
	clubID, err := strconv.ParseUint(clubIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid club ID"})
	}

	isOwner, err := h.service.IsUserClubOwner(c.Request().Context(), uint(clubID), user.ID)
	if err != nil || !isOwner {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "You are not the owner of this club"})
	}

	var req UpdateClubRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	updatedClub, err := h.service.UpdateClub(c.Request().Context(), uint(clubID), req.Name, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedClub)
}
