package attendance

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/go-attendance/internal/database"
	"github.com/umardev500/go-attendance/internal/ent"
	"github.com/umardev500/go-attendance/internal/ent/attendance"
	"github.com/umardev500/go-attendance/internal/ent/user"
)

type Repository interface {
	Create(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Attendance, error)
	GetTodayByUser(ctx context.Context, userID uuid.UUID) (*ent.Attendance, error)
	List(ctx context.Context, params *ListAttendanceParams) ([]*ent.Attendance, int, error)
	Update(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	tm *database.TransactionManager
}

func NewRepository(tm *database.TransactionManager) Repository {
	return &repository{tm: tm}
}

func (r *repository) Create(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error) {
	db := r.tm.FromContext(ctx)
	return db.Attendance.Create().
		SetUsersID(attendance.Edges.Users.ID).
		SetDevicesID(attendance.Edges.Devices.ID).
		Save(ctx)
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Attendance, error) {
	db := r.tm.FromContext(ctx)
	return db.Attendance.Get(ctx, id)
}

func (r *repository) GetTodayByUser(ctx context.Context, userID uuid.UUID) (*ent.Attendance, error) {
	db := r.tm.FromContext(ctx)

	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	return db.Attendance.
		Query().
		Where(
			attendance.HasUsersWith(user.IDEQ(userID)),
			attendance.DateGTE(startOfDay),
			attendance.DateLT(endOfDay),
		).
		First(ctx)
}

func (r *repository) List(ctx context.Context, params *ListAttendanceParams) ([]*ent.Attendance, int, error) {
	db := r.tm.FromContext(ctx)
	q := db.Attendance.Query()

	if params.UserID != nil {
		q = q.Where(attendance.HasUsersWith(user.IDEQ(*params.UserID)))
	}

	if params.Today {
		now := time.Now()
		todayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		tomorrowStart := todayStart.Add(24 * time.Hour)

		q = q.Where(
			attendance.DateGTE(todayStart),
			attendance.DateLT(tomorrowStart),
		)
	}

	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	if params.Limit > 0 {
		q = q.Limit(params.Limit)
	} else {
		q = q.Limit(10)
	}

	if params.Offset > 0 {
		q = q.Offset(params.Offset)
	}

	results, err := q.WithUsers().All(ctx)
	return results, total, err
}

func (r *repository) Update(ctx context.Context, attendance *ent.Attendance) (*ent.Attendance, error) {
	db := r.tm.FromContext(ctx)
	return db.Attendance.UpdateOneID(attendance.ID).
		SetNillableCheckOut(attendance.CheckOut).
		Save(ctx)
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	db := r.tm.FromContext(ctx)
	return db.Attendance.DeleteOneID(id).Exec(ctx)
}
