package repository

import (
	"MathXplains/internal/domain/entity"
	"database/sql"
)

type SubjectRepository struct {
	db *sql.DB
}

func NewSubjectRepository(db *sql.DB) *SubjectRepository {
	return &SubjectRepository{db}
}

func (s *SubjectRepository) Save(subject *domain.Subject) error {

	_, err := s.db.Exec(`INSERT INTO subjects (professor_id, name, full_name)
	VALUES (?, ?, ?);`, subject.ProfessorID, subject.Name, subject.FullName)

	return err
}

func (s *SubjectRepository) FindAll(availableOnly bool) ([]*domain.Subject, error) {
	var condition string
	if availableOnly {
		condition = "WHERE available = 1"
	}

	rows, err := s.db.Query("SELECT * FROM subjects " + condition + ";")
	if err != nil {
		return nil, err
	}

	return serializeSubjects(rows)
}

func (s *SubjectRepository) FindById(id int) (*domain.Subject, error) {
	res := s.db.QueryRow("SELECT * FROM subjects WHERE id = ?", id)

	subj, err := serializeSubject(res)
	if err != nil {
		return nil, err
	}
	return subj, nil
}

func (s *SubjectRepository) DeleteById(id int) error {
	_, err := s.db.Exec("DELETE FROM subjects WHERE id = ?", id)
	return err
}

func serializeSubjects(rows *sql.Rows) ([]*domain.Subject, error) {
	var subjects []*domain.Subject

	for rows.Next() {
		var subject domain.Subject
		err := rows.Scan(
			&subject.ID,
			&subject.ProfessorID,
			&subject.Name,
			&subject.FullName,
			&subject.Available,
		)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, &subject)
	}
	return subjects, nil
}

func serializeSubject(row *sql.Row) (*domain.Subject, error) {
	var subj domain.Subject
	err := row.Scan(
		&subj.ID,
		&subj.ProfessorID,
		&subj.Name,
		&subj.FullName,
		&subj.Available,
	)
	return &subj, err
}
