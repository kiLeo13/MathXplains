package repository

import (
	"MathXplains/internal/domain/entity"
	"MathXplains/internal/domain/sqlite"
	"MathXplains/internal/utils"
	"database/sql"
)

type AppointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{db}
}

// Save saves a new appointment to the database.
//
// For this method, the following fields are ignored and can be safely ommited
// when creating the struct:
//   - ID
//   - ProfessorID
//   - Rejected (defaults to FALSE)
//   - UpdatedAt (defaults to CreatedAt)
//   - Active (defaults to TRUE)
func (a *AppointmentRepository) Save(appt *domain.Appointment) (*domain.Appointment, error) {
	res, err := a.db.Exec(`INSERT INTO appointments (topic, description, user_id, subject_id, scheduled_at, created_at, updated_at, professor_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, (SELECT professor_id FROM subjects WHERE id = ?));`,
		appt.Topic, appt.Description, appt.UserID, appt.SubjectID, appt.ScheduledAt, appt.CreatedAt, appt.CreatedAt, appt.SubjectID)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return a.FindById(int(lastId))
}

func (a *AppointmentRepository) FindAllByUserID(past bool, userId string) ([]*domain.Appointment, error) {
	today := utils.GetTodayf()
	cond := ""
	if past {
		cond = "AND scheduled_at >= " + today
	}

	res, err := a.db.Query("SELECT * FROM appointments WHERE user_id = ? AND active = 1 "+cond+";", userId)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	return deserializeAppointments(res)
}

func (a *AppointmentRepository) FindAll() ([]*domain.Appointment, error) {
	res, err := a.db.Query("SELECT * FROM appointments WHERE active = 1")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	return deserializeAppointments(res)
}

func (a *AppointmentRepository) FindById(id int) (*domain.Appointment, error) {
	res := a.db.QueryRow("SELECT * FROM appointments WHERE id = ? AND active = 1;", id)
	return deserializeAppointment(res)
}

func (a *AppointmentRepository) FindByIdAndUser(userId string, id int) (*domain.Appointment, error) {
	res := a.db.QueryRow("SELECT * FROM appointments WHERE user_id = ? AND id = ? AND active = 1;", userId, id)
	return deserializeAppointment(res)
}

func (a *AppointmentRepository) CountActiveByUserID(userId string) (int, error) {
	today := utils.GetTodayf()
	res := a.db.QueryRow("SELECT COUNT(*) FROM appointments WHERE user_id = ? AND scheduled_at >= ? AND active = 1", userId, today)
	var count int

	err := res.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (a *AppointmentRepository) DeleteById(userid string, id int) error {
	_, err := a.db.Exec("DELETE FROM appointments WHERE id = ? AND user_id = ?", id, userid)
	if err != nil {
		return err
	}
	return nil
}

func deserializeAppointments(rows *sql.Rows) ([]*domain.Appointment, error) {
	var apptms []*domain.Appointment

	for rows.Next() {
		apptm, err := deserializeAppointment(rows)
		if err != nil {
			return nil, err
		}
		apptms = append(apptms, apptm)
	}
	return apptms, nil
}

func deserializeAppointment(row sqlite.RowScanner) (*domain.Appointment, error) {
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
		&apptm.Active,
	)
	if err != nil {
		return nil, err
	}
	return &apptm, nil
}
