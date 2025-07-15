package attendance

import "github.com/google/uuid"

type CheckInRequest struct {
	DeviceID uuid.UUID `json:"device_id" validate:"required"`
	CardUID  string    `json:"card_uid" validate:"required"`
}

type CheckOutRequest struct {
	DeviceID uuid.UUID `json:"device_id"`
	CardUID  string    `json:"card_uid" validate:"required"`
}

type ListAttendanceParams struct {
	Limit  int        `query:"limit"`
	Offset int        `query:"offset"`
	UserID *uuid.UUID `query:"user_id"`
	Today  bool       `query:"today"`
}
