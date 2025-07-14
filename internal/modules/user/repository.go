package user

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/umardev500/go-attendance/internal/database"
	"github.com/umardev500/go-attendance/internal/ent"
	"github.com/umardev500/go-attendance/internal/ent/user"
)

type Repository interface {
	Create(ctx context.Context, user *ent.User) (*ent.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	List(ctx context.Context, params *ListUsersParams) ([]*ent.User, int, error)
	Update(ctx context.Context, user *ent.User) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	tm *database.TransactionManager
}

func NewRepository(tm *database.TransactionManager) Repository {
	return &repository{
		tm: tm,
	}
}

func (r *repository) Create(ctx context.Context, user *ent.User) (*ent.User, error) {
	db := r.tm.FromContext(ctx)
	return db.User.Create().
		SetPhone(*user.Phone).
		SetEmail(*user.Email).
		SetPasswordHash(*user.PasswordHash).
		Save(ctx)
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	db := r.tm.FromContext(ctx)
	return db.User.DeleteOneID(id).Exec(ctx)
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	db := r.tm.FromContext(ctx)
	return db.User.Query().Where(user.EmailEQ(email)).Only(ctx)
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	db := r.tm.FromContext(ctx)
	return db.User.Query().Where(user.IDEQ(id)).Only(ctx)
}

func (r *repository) List(ctx context.Context, params *ListUsersParams) ([]*ent.User, int, error) {
	db := r.tm.FromContext(ctx)
	q := db.User.Query()

	// Filter by search
	if params.Search != "" {
		q = q.Where(user.Or(
			user.EmailContains(params.Search),
			user.PhoneContains(params.Search),
		))
	}

	// Count total before limit & offset
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	// Sorting
	// TODO: add order by field
	if params.SortDir == "desc" {
		q = q.Order(user.ByCreatedAt(sql.OrderDesc()))
	} else {
		q = q.Order(user.ByCreatedAt(sql.OrderAsc()))
	}

	// Limit & Offset
	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	} else {
		q = q.Limit(10)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	users, err := q.All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *repository) Update(ctx context.Context, user *ent.User) (*ent.User, error) {
	db := r.tm.FromContext(ctx)
	return db.User.UpdateOneID(user.ID).
		SetNillablePhone(user.Phone).
		SetNillableEmail(user.Email).
		SetNillablePasswordHash(user.PasswordHash).
		Save(ctx)
}
