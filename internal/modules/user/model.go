package user

import "github.com/google/uuid"

type ListUsersParams struct {
	Limit   int    `query:"limit"`
	Offset  int    `query:"offset"`
	Search  string `query:"search"`
	SortBy  string `query:"sort_by"`
	SortDir string `query:"sort_dir"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"-"`
	Email    *string   `json:"email" validate:"omitempty,email"`
	Password *string   `json:"password"`
	Phone    *string   `json:"phone"`
}
