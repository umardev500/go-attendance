package card

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/umardev500/go-attendance/internal/database"
	"github.com/umardev500/go-attendance/internal/ent"
	"github.com/umardev500/go-attendance/internal/ent/card"
	"github.com/umardev500/go-attendance/internal/ent/user"
)

type Repository interface {
	AssignCardToUser(ctx context.Context, cardID uuid.UUID, userID uuid.UUID) error
	Create(ctx context.Context, card *ent.Card) (*ent.Card, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Card, error)
	GetByUID(ctx context.Context, cardUID string) (*ent.Card, error)
	GetUserByCard(ctx context.Context, cardUID string) (*ent.User, error)
	List(ctx context.Context, params *ListCardParams) ([]*ent.Card, int, error)
	Update(ctx context.Context, card *ent.Card) (*ent.Card, error)
	UnassignCard(ctx context.Context, cardID uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	tm *database.TransactionManager
}

func NewRepository(tm *database.TransactionManager) Repository {
	return &repository{tm: tm}
}

func (r *repository) AssignCardToUser(ctx context.Context, cardID uuid.UUID, userID uuid.UUID) error {
	db := r.tm.FromContext(ctx)
	return db.Card.UpdateOneID(cardID).SetUserID(userID).Exec(ctx)
}

func (r *repository) Create(ctx context.Context, cardData *ent.Card) (*ent.Card, error) {
	db := r.tm.FromContext(ctx)
	return db.Card.Create().
		SetCardUID(*cardData.CardUID).
		SetIssuedAt(cardData.IssuedAt).
		SetIsActive(*cardData.IsActive).
		Save(ctx)
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Card, error) {
	db := r.tm.FromContext(ctx)
	return db.Card.Get(ctx, id)
}

func (r *repository) GetByUID(ctx context.Context, cardUID string) (*ent.Card, error) {
	db := r.tm.FromContext(ctx)
	return db.Card.Query().Where(card.CardUIDEQ(cardUID)).Only(ctx)
}

func (r *repository) GetUserByCard(ctx context.Context, cardUID string) (*ent.User, error) {
	db := r.tm.FromContext(ctx)
	return db.User.
		Query().
		Where(user.HasCardWith(card.CardUIDEQ(cardUID))).
		Only(ctx)
}

func (r *repository) List(ctx context.Context, params *ListCardParams) ([]*ent.Card, int, error) {
	db := r.tm.FromContext(ctx)
	q := db.Card.Query()

	if params.Search != "" {
		q = q.Where(card.CardUIDContains(params.Search))
	}

	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if params.SortDir == "desc" {
		q = q.Order(card.ByIssuedAt(sql.OrderDesc()))
	} else {
		q = q.Order(card.ByIssuedAt(sql.OrderAsc()))
	}

	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	} else {
		q = q.Limit(10)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	cards, err := q.WithUser().All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return cards, total, nil
}

func (r *repository) UnassignCard(ctx context.Context, cardID uuid.UUID) error {
	db := r.tm.FromContext(ctx)
	return db.Card.UpdateOneID(cardID).ClearUser().Exec(ctx)
}

func (r *repository) Update(ctx context.Context, cardData *ent.Card) (*ent.Card, error) {
	db := r.tm.FromContext(ctx)
	return db.Card.UpdateOneID(cardData.ID).
		SetNillableCardUID(cardData.CardUID).
		SetNillableIsActive(cardData.IsActive).
		// SetNillableUsersID((*uuid.UUID)(uuid.Nil)).
		Save(ctx)
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	db := r.tm.FromContext(ctx)
	return db.Card.DeleteOneID(id).Exec(ctx)
}
