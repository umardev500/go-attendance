package device

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
	router.Post("/", h.Create)
	router.Put("/:id", h.Update)
	router.Delete("/:id", h.Delete)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateDeviceRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to create device", err))
	}

	device, err := h.service.Create(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to create device", err))
	}

	return c.JSON(api.Success(device, "Device created successfully"))
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Invalid device ID", err))
	}

	device, err := h.service.GetByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to get device", err))
	}

	return c.JSON(api.Success(device, "Device retrieved successfully"))
}

func (h *Handler) List(c *fiber.Ctx) error {
	params := &ListDeviceParams{}
	if err := c.QueryParser(params); err != nil {
		return err
	}

	devices, total, err := h.service.List(c.UserContext(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to get devices", err))
	}

	resp := api.PaginatedSuccess(devices, params.Limit, params.Offset, total, "Devices retrieved successfully")

	return c.JSON(resp)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Invalid device ID", err))
	}

	var req UpdateDeviceRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to update device", err))
	}

	req.ID = id

	device, err := h.service.Update(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to update device", err))
	}

	return c.JSON(api.Success(device, "Device updated successfully"))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Invalid device ID", err))
	}

	if err := h.service.Delete(c.UserContext(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to delete device", err))
	}

	return c.JSON(api.Success(nil, "Device deleted successfully"))
}
