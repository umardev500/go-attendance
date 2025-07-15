package attendance

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/umardev500/go-attendance/pkg/api"
)

type Handler struct {
	service  Service
	validate *validator.Validate
}

func NewHandler(service Service, validate *validator.Validate) *Handler {
	return &Handler{
		service:  service,
		validate: validate,
	}
}

func (h *Handler) Setup(router fiber.Router) {
	router.Get("/", h.List)
	router.Get("/:id", h.GetByID)
	router.Post("/check-in", h.CheckIn)
	router.Put("/check-out", h.CheckOut)
}

func (h *Handler) CheckIn(c *fiber.Ctx) error {
	var req CheckInRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to check-in", err))
	}

	attendance, err := h.service.CheckIn(c.UserContext(), &req)
	if err != nil {
		if err == ErrAlreadyCheckedIn {
			return c.Status(fiber.StatusConflict).JSON(api.Error("User already checked in", err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to check-in", err))
	}

	return c.JSON(api.Success(attendance, "Check-in successful"))
}

func (h *Handler) CheckOut(c *fiber.Ctx) error {
	var req CheckOutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Invalid request", err))
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to check-out", err))
	}

	attendance, err := h.service.CheckOut(c.UserContext(), &req)
	if err != nil {
		if err == ErrAlreadyCheckedOut {
			return c.Status(fiber.StatusConflict).JSON(api.Error("User already checked out", err))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to check-out", err))
	}

	return c.JSON(api.Success(attendance, "Check-out successful"))
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid UUID")
	}

	attendance, err := h.service.GetByID(c.UserContext(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(api.Success(attendance, "Attendance retrieved successfully"))
}

func (h *Handler) List(c *fiber.Ctx) error {
	params := &ListAttendanceParams{}
	if err := c.QueryParser(params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid query params")
	}

	attendances, total, err := h.service.List(c.UserContext(), params)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	resp := api.PaginatedSuccess(attendances, params.Limit, params.Offset, total, "Attendances retrieved successfully")
	return c.JSON(resp)
}
