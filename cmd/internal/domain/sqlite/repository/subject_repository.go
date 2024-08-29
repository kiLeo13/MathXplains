package repository

import (
	"MathXplains/internal/domain/entity"
	"MathXplains/internal/domain/sqlite"
	"database/sql"
)

type SubjectRepository struct {
	db *sql.DB
}

func NewSubjectRepository(db *sql.DB) *SubjectRepository {
	return &SubjectRepository{db}
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
	return deserializeSubjects(rows)
}

func (s *SubjectRepository) FindById(id int) (*domain.Subject, error) {
	res := s.db.QueryRow("SELECT * FROM subjects WHERE id = ?", id)

	subj, err := deserializeSubject(res)
	if err != nil {
		return nil, err
	}
	return subj, nil
}

func (s *SubjectRepository) DeleteById(id int) error {
	_, err := s.db.Exec("DELETE FROM subjects WHERE id = ?", id)
	return err
}

func deserializeSubjects(rows *sql.Rows) ([]*domain.Subject, error) {
	var subjects []*domain.Subject

	for rows.Next() {
		subj, err := deserializeSubject(rows)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subj)
	}
	return subjects, nil
}

func deserializeSubject(row sqlite.RowScanner) (*domain.Subject, error) {
	var subj domain.Subject
	err := row.Scan(
		&subj.ID,
		&subj.ProfessorID,
		&subj.Name,
		&subj.FullName,
		&subj.Available,
	)
	if err != nil {
		return nil, err
	}
	return &subj, nil
}
