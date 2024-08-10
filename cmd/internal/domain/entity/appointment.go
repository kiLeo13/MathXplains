package domain

import (
	"MathXplains/internal/utils"
	"time"
)

type Appointment struct {
	ID          int
	Topic       string
	Description string
	UserID      string
	SubjectID   int
	ProfessorID *int
	Rejected    bool
	ScheduledAt int64
	CreatedAt   int64
	UpdatedAt   int64
}

func (a *Appointment) IsActive() bool {
	today := utils.GetTodayMidnight()
	// At this point, we are safe to assume that the ScheduledAt field
	// will always point to (00:00:00 UTC)
	schd := time.Unix(a.ScheduledAt, 0)
	return schd.Unix() >= today
}
