package entity

import (
	"github.com/google/uuid"
)

type Bookings struct {
	ID               uuid.UUID `json:"id,omitempty"`
	DoctorScheduleID uuid.UUID `json:"doctor_schedule_id,omitempty"`
	PatientID        uuid.UUID `json:"patient_id,omitempty"`
	BookingDate             string    `json:"booking_date,omitempty"`
	MstScheduleID    int       `json:"number,omitempty"` //refer to mst_schedule id
	Complaint        string    `json:"complaint,omitempty"`
	Status           string    `json:"status,omitempty"` //(waiting, done, canceled)
	CreatedAt        string    `json:"created_at,omitempty"`
	UpdatedAt        string    `json:"updated_at,omitempty"`
	DeletedAt        string    `json:"deleted_at,omitempty"`
}
