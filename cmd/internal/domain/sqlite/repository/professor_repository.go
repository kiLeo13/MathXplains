package repository

import (
	"MathXplains/internal/domain/entity"
	"MathXplains/internal/domain/sqlite"
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
	return deserializeProfessors(res)
}

func (p *ProfessorRepository) DeleteById(id int) error {
	_, err := p.db.Exec("DELETE FROM professors WHERE id = ?;", id)
	return err
}

func deserializeProfessors(rows *sql.Rows) ([]*domain.Professor, error) {
	var professors []*domain.Professor

	for rows.Next() {
		prof, err := deserializeProfessor(rows)
		if err != nil {
			return nil, err
		}
		professors = append(professors, prof)
	}
	return professors, nil
}

func deserializeProfessor(row sqlite.RowScanner) (*domain.Professor, error) {
	var prof domain.Professor

	err := row.Scan(
		&prof.ID,
		&prof.Name,
		&prof.FullName,
		&prof.Nick,
		&prof.Known,
	)
	if err != nil {
		return nil, err
	}
	return &prof, nil
}
