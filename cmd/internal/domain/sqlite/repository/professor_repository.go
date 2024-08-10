package repository

import (
	"MathXplains/internal/domain/entity"
	"database/sql"
)

type ProfessorRepository struct {
	db *sql.DB
}

func NewProfessorRepository(db *sql.DB) *ProfessorRepository {
	return &ProfessorRepository{db}
}

func (p *ProfessorRepository) FindAll(knownOnly bool) ([]*domain.Professor, error) {

	var condition string
	if knownOnly {
		condition = "WHERE known = 1"
	}

	res, err := p.db.Query("SELECT * FROM professors " + condition + ";")
	if err != nil {
		return nil, err
	}

	return serializeProfessors(res)
}

func (p *ProfessorRepository) DeleteById(id int) error {

	_, err := p.db.Exec("DELETE FROM professors WHERE id = ?;", id)
	return err
}

func serializeProfessors(rows *sql.Rows) ([]*domain.Professor, error) {
	var professors []*domain.Professor

	for rows.Next() {
		var professor domain.Professor
		err := rows.Scan(
			&professor.ID,
			&professor.Name,
			&professor.FullName,
			&professor.Nick,
			&professor.Known,
		)
		if err != nil {
			return nil, err
		}
		professors = append(professors, &professor)
	}
	return professors, nil
}
