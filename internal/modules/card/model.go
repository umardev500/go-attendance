package card

import (
	"time"

	"github.com/google/uuid"
)

type CreateCardRequest struct {
	CardUID  string    `json:"card_uid" validate:"required"`
	IssuedAt time.Time `json:"issued_at"`
	IsActive bool      `json:"is_active"`
}

type AssignCardToUser struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type UpdateCardRequest struct {
	ID       uuid.UUID         `json:"-"`
	CardUID  *string           `json:"card_uid"`
	IsActive *bool             `json:"is_active"`
	IssuedAt time.Time         `json:"-"`
	Assign   *AssignCardToUser `json:"assign"`
	Unassign bool              `json:"unassign"`
}

type ListCardParams struct {
	Limit   int    `query:"limit"`
	Offset  int    `query:"offset"`
	Search  string `query:"search"`
	SortDir string `query:"sort_dir"`
}
