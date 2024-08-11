package service

import (
	domain "MathXplains/internal/domain/entity"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	MaxActiveAppointments = 20
	// maxDeletionPeriod is the maximum amount of time
	// a user can wait before deleting an appointment
	maxDeletionPeriod = 24 * time.Hour

	minTopicLength = 5
	maxTopicLength = 30

	minDescLength = 10
	maxDescLength = 1000
)

type AppointmentDTO struct {
	ID          int    `json:"id"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	SubjectID   int    `json:"subject_id"`
	ProfessorID *int   `json:"professor_id"`
	Rejected    bool   `json:"rejected"`
	IsActive    bool   `json:"is_active"`
	ScheduledAt string `json:"scheduled_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type AppointmentCreateDTO struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
	ScheduledAt string `json:"scheduled_at"`
	SubjectID   int    `json:"subject_id"`
}

func CreateAppointment(userId string, data *AppointmentCreateDTO) (*AppointmentDTO, *APIError) {
	if err := checkCountLimit(userId); err != nil {
		return nil, err
	}

	topic := strings.TrimSpace(data.Topic)
	desc := strings.TrimSpace(data.Description)

	if err := checkTopic(topic); err != nil {
		return nil, err
	}

	if err := checkDescription(desc); err != nil {
		return nil, err
	}

	if err := checkSubject(data.SubjectID); err != nil {
		return nil, err
	}

	scheduledAt, err := ToEpoch(time.DateOnly, data.ScheduledAt)
	if err != nil {
		return nil, ErrorIncorrectDateFormat
	}

	if isDateInPast(scheduledAt) {
		return nil, ErrorDateInThePast
	}

	now := NowUTC()
	newApptm, err := apptmRepo.Save(topic, desc, userId, data.SubjectID, scheduledAt, now)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	return toApptmDTO(newApptm), nil
}

func GetAppointments(all bool, status, userId string) ([]*AppointmentDTO, *APIError) {
	status = strings.ToLower(status)
	apptmts, err := getFiltered(all, userId)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	apptmList := make([]*AppointmentDTO, 0)
	for _, a := range apptmts {
		apptmList = append(apptmList, toApptmDTO(a))
	}
	return apptmList, nil
}

func DeleteAppointment(userId, id string) *APIError {
	apptId, err := strconv.Atoi(id)
	if err != nil {
		return ErrorInvalidPathParam("id", "integer")
	}

	apptm, err := apptmRepo.Find(userId, apptId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
		return ErrorInternalServer
	}

	if apptm == nil {
		return nil
	}

	if err := checkDeletionDate(apptm.CreatedAt); err != nil {
		return err
	}

	err = apptmRepo.DeleteById(userId, apptId)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func getFiltered(all bool, userId string) ([]*domain.Appointment, error) {
	if all {
		return apptmRepo.FindAll()
	} else {
		return apptmRepo.FindByUserID(userId)
	}
}

func checkCountLimit(userId string) *APIError {
	count, err := apptmRepo.CountActiveByUserID(userId)
	if err != nil {
		fmt.Println(err)
		return ErrorInternalServer
	}

	if count > MaxActiveAppointments {
		return ErrorTooManyAppointments
	}
	return nil
}

func checkTopic(topic string) *APIError {
	if len(topic) < minTopicLength || len(topic) > maxTopicLength {
		return ErrorInvalidTopicRange
	}
	return nil
}

func checkDescription(description string) *APIError {
	if len(description) < minDescLength || len(description) > maxDescLength {
		return ErrorInvalidDescriptionRange
	}
	return nil
}

func checkDeletionDate(sec int64) *APIError {
	schd := time.Unix(sec, 0)
	now := time.Now().UTC()
	period := now.Sub(schd)

	if period >= maxDeletionPeriod {
		return ErrorDateInThePast
	}
	return nil
}

func checkSubject(subjectId int) *APIError {
	subj, err := subjectRepo.FindById(subjectId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
		return ErrorInternalServer
	}

	if subj == nil {
		return ErrorSubjectDoesNotExist
	}

	if subj.Available != 1 {
		return ErrorSubjectUnavailable
	}

	return nil
}

func toApptmDTO(a *domain.Appointment) *AppointmentDTO {
	return &AppointmentDTO{
		ID:          a.ID,
		Topic:       a.Topic,
		Description: a.Description,
		UserID:      a.UserID,
		SubjectID:   a.SubjectID,
		ProfessorID: a.ProfessorID,
		Rejected:    a.Rejected,
		IsActive:    a.IsActive(),
		ScheduledAt: FormatDate(a.ScheduledAt),
		CreatedAt:   FormatEpoch(a.CreatedAt),
		UpdatedAt:   FormatEpoch(a.UpdatedAt),
	}
}
