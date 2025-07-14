package device

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/umardev500/go-attendance/internal/database"
	"github.com/umardev500/go-attendance/internal/ent"
	"github.com/umardev500/go-attendance/internal/ent/device"
)

type Repository interface {
	Create(ctx context.Context, device *ent.Device) (*ent.Device, error)
	GetByID(ctx context.Context, id int) (*ent.Device, error)
	List(ctx context.Context, params *ListDeviceParams) ([]*ent.Device, int, error)
	Update(ctx context.Context, device *ent.Device) (*ent.Device, error)
	Delete(ctx context.Context, id int) error
}

type repository struct {
	tm *database.TransactionManager
}

func NewRepository(tm *database.TransactionManager) Repository {
	return &repository{
		tm: tm,
	}
}

func (r *repository) Create(ctx context.Context, d *ent.Device) (*ent.Device, error) {
	db := r.tm.FromContext(ctx)
	return db.Device.Create().
		SetName(*d.Name).
		SetLocation(*d.Location).
		SetInstalledAt(*d.InstalledAt).
		SetIsActive(*d.IsActive).
		Save(ctx)
}

func (r *repository) GetByID(ctx context.Context, id int) (*ent.Device, error) {
	db := r.tm.FromContext(ctx)
	return db.Device.Get(ctx, id)
}

func (r *repository) List(ctx context.Context, params *ListDeviceParams) ([]*ent.Device, int, error) {
	db := r.tm.FromContext(ctx)
	q := db.Device.Query()

	// Filter by search (name or location)
	if params.Search != "" {
		q = q.Where(
			device.Or(
				device.NameContains(params.Search),
				device.LocationContains(params.Search),
			),
		)
	}

	// Count total before applying limit/offset
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Sorting by installed_at
	if params.SortDir == "desc" {
		q = q.Order(device.ByInstalledAt(sql.OrderDesc()))
	} else {
		q = q.Order(device.ByInstalledAt(sql.OrderAsc()))
	}

	// Limit & offset
	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	} else {
		q = q.Limit(10)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	devices, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return devices, total, nil
}

func (r *repository) Update(ctx context.Context, d *ent.Device) (*ent.Device, error) {
	db := r.tm.FromContext(ctx)
	return db.Device.UpdateOneID(d.ID).
		SetNillableName(d.Name).
		SetNillableLocation(d.Location).
		SetNillableInstalledAt(d.InstalledAt).
		SetNillableIsActive(d.IsActive).
		Save(ctx)
}

func (r *repository) Delete(ctx context.Context, id int) error {
	db := r.tm.FromContext(ctx)
	return db.Device.DeleteOneID(id).Exec(ctx)
}
