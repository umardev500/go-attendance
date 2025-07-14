package device

import (
	"context"

	"github.com/umardev500/go-attendance/internal/ent"
)

type Service interface {
	Create(ctx context.Context, req *CreateDeviceRequest) (*ent.Device, error)
	GetByID(ctx context.Context, id int) (*ent.Device, error)
	List(ctx context.Context, params *ListDeviceParams) ([]*ent.Device, int, error)
	Update(ctx context.Context, req *UpdateDeviceRequest) (*ent.Device, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, req *CreateDeviceRequest) (*ent.Device, error) {
	device := &ent.Device{
		Name:        &req.Name,
		Location:    &req.Location,
		InstalledAt: &req.InstalledAt,
		IsActive:    &req.IsActive,
	}

	return s.repo.Create(ctx, device)
}

func (s *service) GetByID(ctx context.Context, id int) (*ent.Device, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) List(ctx context.Context, params *ListDeviceParams) ([]*ent.Device, int, error) {
	return s.repo.List(ctx, params)
}

func (s *service) Update(ctx context.Context, req *UpdateDeviceRequest) (*ent.Device, error) {
	device := &ent.Device{
		ID:          req.ID,
		Name:        req.Name,
		Location:    req.Location,
		InstalledAt: req.InstalledAt,
		IsActive:    req.IsActive,
	}

	return s.repo.Update(ctx, device)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
