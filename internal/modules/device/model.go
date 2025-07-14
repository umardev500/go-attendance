package device

import "time"

type ListDeviceParams struct {
	Limit   int    `query:"limit"`
	Offset  int    `query:"offset"`
	Search  string `query:"search"`
	SortBy  string `query:"sort_by"`
	SortDir string `query:"sort_dir"`
}

type CreateDeviceRequest struct {
	Name        string    `json:"name" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	InstalledAt time.Time `json:"installed_at"`
	IsActive    bool      `json:"is_active"`
}

type UpdateDeviceRequest struct {
	ID          int        `json:"-"`
	Name        *string    `json:"name" validate:"required"`
	Location    *string    `json:"location" validate:"required"`
	InstalledAt *time.Time `json:"installed_at"`
	IsActive    *bool      `json:"is_active"`
}
