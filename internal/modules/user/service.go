package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-attendance/internal/ent"
	"github.com/umardev500/go-attendance/pkg/crypto"
)

type Service interface {
	Create(ctx context.Context, req *CreateUserRequest) (*ent.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
	GetByEmail(ctx context.Context, email string) (*ent.User, error)
	List(ctx context.Context, params *ListUsersParams) ([]*ent.User, int, error)
	Update(ctx context.Context, req *UpdateUserRequest) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, req *CreateUserRequest) (*ent.User, error) {
	passwordHash, err := crypto.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &ent.User{
		Phone:        &req.Phone,
		Email:        &req.Email,
		PasswordHash: &passwordHash,
	}

	return s.repo.Create(ctx, user)
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *service) GetByEmail(ctx context.Context, email string) (*ent.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) List(ctx context.Context, params *ListUsersParams) ([]*ent.User, int, error) {
	users, total, err := s.repo.List(ctx, params)
	return users, total, err
}

func (s *service) Update(ctx context.Context, req *UpdateUserRequest) (*ent.User, error) {
	var passwordHash *string
	if req.Password != nil {
		hashedPassword, err := crypto.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
		passwordHash = &hashedPassword
	}

	user := &ent.User{
		ID:           req.ID,
		Phone:        req.Phone,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	s.repo.Update(ctx, user)
	return s.repo.GetByID(ctx, req.ID)
}
