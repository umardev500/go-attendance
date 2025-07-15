package card

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/go-attendance/internal/database"
	"github.com/umardev500/go-attendance/internal/ent"
)

type Service interface {
	Create(ctx context.Context, req *CreateCardRequest) (*ent.Card, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Card, error)
	List(ctx context.Context, params *ListCardParams) ([]*ent.Card, int, error)
	Update(ctx context.Context, req *UpdateCardRequest) (*ent.Card, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
	tm   *database.TransactionManager
}

func NewService(repo Repository, tm *database.TransactionManager) Service {
	return &service{
		repo: repo,
		tm:   tm,
	}
}

func (s *service) Create(ctx context.Context, req *CreateCardRequest) (*ent.Card, error) {
	cardData := &ent.Card{
		CardUID:  &req.CardUID,
		IssuedAt: req.IssuedAt,
		IsActive: &req.IsActive,
	}

	return s.repo.Create(ctx, cardData)
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*ent.Card, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) List(ctx context.Context, params *ListCardParams) ([]*ent.Card, int, error) {
	return s.repo.List(ctx, params)
}

func (s *service) Update(ctx context.Context, req *UpdateCardRequest) (*ent.Card, error) {
	var c *ent.Card

	err := s.tm.WithTx(ctx, func(ctx context.Context) error {
		cardData := &ent.Card{
			ID:       req.ID,
			CardUID:  req.CardUID,
			IsActive: req.IsActive,
		}

		var err error

		c, err = s.repo.Update(ctx, cardData)
		if err != nil {
			return err
		}

		if req.Unassign {
			err = s.repo.UnassignCard(ctx, req.ID)
			if err != nil {
				return err
			}
		} else if req.Assign != nil {
			err = s.repo.AssignCardToUser(ctx, req.ID, req.Assign.UserID)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
