package attendance

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-attendance/internal/ent"
	"github.com/umardev500/go-attendance/internal/modules/card"
)

type Service interface {
	CheckIn(ctx context.Context, req *CheckInRequest) (*ent.Attendance, error)
	CheckOut(ctx context.Context, req *CheckOutRequest) (*ent.Attendance, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Attendance, error)
	List(ctx context.Context, params *ListAttendanceParams) ([]*ent.Attendance, int, error)
}

type service struct {
	repo     Repository
	cardRepo card.Repository
}

func NewService(repo Repository, cardRepo card.Repository) Service {
	return &service{
		repo:     repo,
		cardRepo: cardRepo,
	}
}

func (s *service) CheckIn(ctx context.Context, req *CheckInRequest) (*ent.Attendance, error) {
	user, err := s.cardRepo.GetUserByCard(ctx, req.CardUID)
	if err != nil {
		return nil, err
	}

	existing, err := s.repo.GetTodayByUser(ctx, user.ID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existing != nil {
		return nil, ErrAlreadyCheckedIn
	}

	att := &ent.Attendance{
		Edges: ent.AttendanceEdges{
			Users:   &ent.User{ID: user.ID},
			Devices: &ent.Device{ID: req.DeviceID},
		},
	}

	return s.repo.Create(ctx, att)
}

func (s *service) CheckOut(ctx context.Context, req *CheckOutRequest) (*ent.Attendance, error) {
	user, err := s.cardRepo.GetUserByCard(ctx, req.CardUID)
	if err != nil {
		return nil, err
	}

	attendance, err := s.repo.GetTodayByUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	if attendance.CheckOut != nil {
		return nil, ErrAlreadyCheckedOut
	}

	now := time.Now()
	attendance.CheckOut = &now

	return s.repo.Update(ctx, attendance)
}

func (s *service) GetByID(ctx context.Context, id uuid.UUID) (*ent.Attendance, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) List(ctx context.Context, params *ListAttendanceParams) ([]*ent.Attendance, int, error) {
	return s.repo.List(ctx, params)
}
