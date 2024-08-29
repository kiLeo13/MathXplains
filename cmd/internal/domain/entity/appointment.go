package domain

import (
	"MathXplains/internal/utils"
	"fmt"
)

type Appointment struct {
	ID          int
	Topic       string
	Description string
	UserID      string
	SubjectID   int
	ProfessorID *int
	Rejected    bool
	ScheduledAt string
	CreatedAt   int64
	UpdatedAt   int64
	Active      bool
}

// IsActive does not have any relation to the "active" soft deletion column in the database,
// instead, it returns whether the ScheduledAt is (x >= today) or not.
func (a *Appointment) IsActive() bool {
	today := utils.GetToday()
	schd, err := utils.DateFromFormat(a.ScheduledAt)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return !schd.Before(*today)
}
