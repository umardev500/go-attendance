package card

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
	var req CreateCardRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to create card", err))
	}

	card, err := h.service.Create(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to create card", err))
	}

	return c.JSON(api.Success(card, "Card created successfully"))
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Invalid card ID", err))
	}

	card, err := h.service.GetByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to get card", err))
	}

	return c.JSON(api.Success(card, "Card retrieved successfully"))
}

func (h *Handler) List(c *fiber.Ctx) error {
	params := &ListCardParams{}
	if err := c.QueryParser(params); err != nil {
		return err
	}

	cards, total, err := h.service.List(c.UserContext(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to get cards", err))
	}

	resp := api.PaginatedSuccess(cards, params.Limit, params.Offset, total, "Cards retrieved successfully")

	return c.JSON(resp)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Invalid card ID", err))
	}

	var req UpdateCardRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to update card", err))
	}

	req.ID = id

	card, err := h.service.Update(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to update card", err))
	}

	return c.JSON(api.Success(card, "Card updated successfully"))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Invalid card ID", err))
	}

	if err := h.service.Delete(c.UserContext(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to delete card", err))
	}

	return c.JSON(api.Success(nil, "Card deleted successfully"))
}
