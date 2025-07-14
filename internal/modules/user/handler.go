package user

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
	router.Delete("/:id", h.Delete)
	router.Put("/:id", h.Update)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to create user", err))
	}

	user, err := h.service.Create(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to create user", err))
	}

	return c.JSON(api.Success(user, "User created successfully"))
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to delete user", err))
	}

	if err := h.service.Delete(c.UserContext(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to delete user", err))
	}

	return c.JSON(api.Success(nil, "User deleted successfully"))
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to get user", err))
	}

	user, err := h.service.GetByID(c.UserContext(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to get user", err))
	}

	return c.JSON(api.Success(user, "User retrieved successfully"))
}

func (h *Handler) List(c *fiber.Ctx) error {
	params := &ListUsersParams{}
	if err := c.QueryParser(params); err != nil {
		return err
	}

	users, total, err := h.service.List(c.UserContext(), params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to get users", err))
	}

	resp := api.PaginatedSuccess(users, params.Limit, params.Offset, total, "Get all users")

	return c.JSON(resp)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to update user", err))
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(api.Error("Failed to update user", err))
	}

	req.ID = id

	user, err := h.service.Update(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(api.Error("Failed to update user", err))
	}

	return c.JSON(api.Success(user, "User updated successfully"))
}
