// Package repository is a crap
package repository

import (
	"MathXplains/internal/domain/entity"
	"MathXplains/internal/utils"
	"database/sql"
)

type AppointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{db}
}

func (a *AppointmentRepository) Save(topic, description, userId string, subjectId int, scheduled, timestamp int64) (*domain.Appointment, error) {
	res, err := a.db.Exec(`INSERT INTO appointments (topic, description, user_id, subject_id, scheduled_at, created_at, updated_at, professor_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, (SELECT professor_id FROM subjects WHERE id = ?));`,
		topic, description, userId, subjectId, scheduled, timestamp, timestamp, subjectId)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return a.FindById(int(lastId))
}

func (a *AppointmentRepository) FindByUserID(userId string) ([]*domain.Appointment, error) {
	res, err := a.db.Query("SELECT * FROM appointments WHERE user_id = ?;", userId)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	return serializeAppointments(res)
}

func (a *AppointmentRepository) FindAll() ([]*domain.Appointment, error) {
	res, err := a.db.Query("SELECT * FROM appointments")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	return serializeAppointments(res)
}

func (a *AppointmentRepository) FindById(id int) (*domain.Appointment, error) {
	res := a.db.QueryRow("SELECT * FROM appointments WHERE id = ?;", id)
	return serializeAppointment(res)
}

func (a *AppointmentRepository) Find(userId string, id int) (*domain.Appointment, error) {
	res := a.db.QueryRow("SELECT * FROM appointments WHERE user_id = ? AND id = ?;", userId, id)
	return serializeAppointment(res)
}

func (a *AppointmentRepository) CountActiveByUserID(userId string) (int, error) {
	today := utils.GetTodayMidnight()
	res := a.db.QueryRow("SELECT COUNT(*) FROM appointments WHERE user_id = ? AND scheduled_at >= ?", userId, today)
	var count int

	err := res.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *AppointmentRepository) DeleteById(userid string, id int) error {
	_, err := a.db.Exec("DELETE FROM appointments WHERE id = ? AND user_id = ?", id, userid)
	return err
}

func serializeAppointments(rows *sql.Rows) ([]*domain.Appointment, error) {
	var apptms []*domain.Appointment

	for rows.Next() {
		var apptm domain.Appointment
		err := rows.Scan(
			&apptm.ID,
			&apptm.Topic,
			&apptm.Description,
			&apptm.UserID,
			&apptm.SubjectID,
			&apptm.ProfessorID,
			&apptm.Rejected,
			&apptm.ScheduledAt,
			&apptm.CreatedAt,
			&apptm.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		apptms = append(apptms, &apptm)
	}
	return apptms, nil
}

func serializeAppointment(row *sql.Row) (*domain.Appointment, error) {
	var apptm domain.Appointment

	err := row.Scan(
		&apptm.ID,
		&apptm.Topic,
		&apptm.Description,
		&apptm.UserID,
		&apptm.SubjectID,
		&apptm.ProfessorID,
		&apptm.Rejected,
		&apptm.ScheduledAt,
		&apptm.CreatedAt,
		&apptm.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &apptm, nil
}
