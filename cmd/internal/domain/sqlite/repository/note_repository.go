package repository

import (
	"MathXplains/internal/domain/entity"
	"MathXplains/internal/domain/sqlite"
	"database/sql"
	"fmt"
)

type NoteRepository struct {
	db *sql.DB
}

func NewNoteRepository(db *sql.DB) *NoteRepository {
	return &NoteRepository{db}
}

// Save saves a new note entity to the database, the ID column will be ignored.
func (n *NoteRepository) Save(note *domain.Note) (*domain.Note, error) {
	res, err := n.db.Exec(`INSERT INTO notes (profile, name, content, last_modified)
		VALUES (?, ?, ?, ?);`, note.Profile, note.Name, note.Content, note.LastModified)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Failed to get last id at saving %v\n", note)
		return nil, err
	}
	return n.FindByID(int(lastId))
}

// Update performs an update to the note for the given ID,
// overwriting all the values with the ones in the provided struct.
func (n *NoteRepository) Update(note *domain.Note) (*domain.Note, error) {
	_, err := n.db.Exec(`UPDATE notes
		SET profile = ?, name = ?, content = ?, last_modified = ? WHERE id = ? AND active = 1`,
		note.Profile, note.Name, note.Content, note.LastModified, note.ID)
	if err != nil {
		return nil, err
	}
	return n.FindByID(note.ID)
}

func (n *NoteRepository) FindByID(id int) (*domain.Note, error) {
	res := n.db.QueryRow("SELECT * FROM notes WHERE id = ? AND active = 1", id)
	note, err := deserializeNote(res)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n *NoteRepository) ExistsByProfileAndName(profile, name string) (bool, error) {
	res := n.db.QueryRow("SELECT EXISTS(SELECT * FROM notes WHERE profile = ? AND name = ? AND active = 1)", profile, name)

	var value int
	err := res.Scan(&value)
	if err != nil {
		return false, err
	}
	return value == 1, nil
}

func (n *NoteRepository) FindAllByProfile(profile string) ([]*domain.Note, error) {
	res, err := n.db.Query("SELECT * FROM notes WHERE profile = ? AND active = 1;", profile)
	if err != nil {
		return nil, err
	}
	return deserializeNotes(res)
}

// DeleteByID soft-deletes a row for the given id.
func (n *NoteRepository) DeleteByID(id int) error {
	_, err := n.db.Exec("UPDATE notes SET active = 0 WHERE id = ?;", id)
	if err != nil {
		return err
	}
	return nil
}

func deserializeNotes(rows *sql.Rows) ([]*domain.Note, error) {
	var notes []*domain.Note

	for rows.Next() {
		note, err := deserializeNote(rows)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

func deserializeNote(row sqlite.RowScanner) (*domain.Note, error) {
	var note domain.Note

	err := row.Scan(
		&note.ID,
		&note.Profile,
		&note.Name,
		&note.Content,
		&note.LastModified,
		&note.Active,
	)
	if err != nil {
		return nil, err
	}
	return &note, nil
}
