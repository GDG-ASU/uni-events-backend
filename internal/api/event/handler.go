package event

import (
	"net/http"
	"strconv"
	"uni-events-backend/internal/models"
	"uni-events-backend/internal/service"
	"uni-events-backend/pkg/utils"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	clubService   service.ClubService
	userService   service.UserService
	eventService  service.EventService
}

func NewHandler(clubService service.ClubService, userService service.UserService, eventService service.EventService) *Handler {
	return &Handler{
		clubService:  clubService,
		userService:  userService,
		eventService: eventService,
	}
}

type CreateEventRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	ClubID      uint   `json:"club_id"`
}

type UpdateEventRequest struct {
	EventID     uint   `json:"event_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func (h *Handler) CreateEvent(c echo.Context) error {
	clerkID := utils.GetClerkUserID(c)

	user, err := h.userService.GetUserByClerkID(c.Request().Context(), clerkID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
	}

	var req CreateEventRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	isOwner, err := h.clubService.IsUserClubOwner(c.Request().Context(), req.ClubID, user.ID)
	if err != nil || !isOwner {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "You are not the owner of this club"})
	}

	event := &models.Event{
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
		ClubID:      req.ClubID,
	}

	createdEvent, err := h.eventService.CreateEvent(c.Request().Context(), event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdEvent)
}

func (h *Handler) ListEvents(c echo.Context) error {
	events, err := h.eventService.GetAllEvents(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch events"})
	}
	return c.JSON(http.StatusOK, events)
}

func (h *Handler) UpdateEvent(c echo.Context) error {
	clerkID := utils.GetClerkUserID(c)

	user, err := h.userService.GetUserByClerkID(c.Request().Context(), clerkID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
	}

	eventIDParam := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid event ID"})
	}

	existingEvent, err := h.eventService.GetEventByID(c.Request().Context(), uint(eventID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Event not found"})
	}

	isOwner, err := h.clubService.IsUserClubOwner(c.Request().Context(), existingEvent.ClubID, user.ID)
	if err != nil || !isOwner {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "You are not the owner of this club"})
	}

	var req CreateEventRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	existingEvent.Title = req.Title
	existingEvent.Description = req.Description
	existingEvent.Date = req.Date

	updatedEvent, err := h.eventService.UpdateEvent(c.Request().Context(), existingEvent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedEvent)
}


func (h *Handler) DeleteEvent(c echo.Context) error {
	clerkID := utils.GetClerkUserID(c)

	user, err := h.userService.GetUserByClerkID(c.Request().Context(), clerkID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Unauthorized"})
	}

	eventIDParam := c.Param("id")
	eventID, err := strconv.ParseUint(eventIDParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid event ID"})
	}

	event, err := h.eventService.GetEventByID(c.Request().Context(), uint(eventID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Event not found"})
	}

	isOwner, err := h.clubService.IsUserClubOwner(c.Request().Context(), event.ClubID, user.ID)
	if err != nil || !isOwner {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "You are not the owner of this club"})
	}

	deletedEvent, err := h.eventService.DeleteEvent(c.Request().Context(), event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Event deleted successfully",
		"event":   deletedEvent,
	})
}
