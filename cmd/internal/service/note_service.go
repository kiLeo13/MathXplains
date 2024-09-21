package service

import (
	domain "MathXplains/internal/domain/entity"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

const (
	noteNameMinLength = 1
	noteNameMaxLength = 30

	noteProfileMinlength = 3
	noteProfileMaxlength = 50

	noteContentMaxLength = 5_000_000
)

type NoteDTO struct {
	ID           int    `json:"id"`
	Profile      string `json:"profile"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	LastModified string `json:"last_modified"`
}

type NoteCreateDTO struct {
	Profile string `json:"profile"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func GetNotes(author string) ([]*NoteDTO, *APIError) {
	notes, err := noteRepo.FindAllByProfile(author)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	noteList := make([]*NoteDTO, len(notes))
	for i, n := range notes {
		noteList[i] = toNoteDTO(n)
	}
	return noteList, nil
}

func CreateNote(note *NoteCreateDTO) (*NoteDTO, *APIError) {
	note.Name = strings.TrimSpace(note.Name)
	note.Profile = strings.TrimSpace(note.Profile)
	found, err := noteRepo.ExistsByProfileAndName(note.Profile, note.Name)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}
	if found {
		return nil, ErrorNoteNameAlreadyExists
	}

	if err := checkNoteName(note.Name); err != nil {
		return nil, err
	}

	if err := checkNoteProfile(note.Profile); err != nil {
		return nil, err
	}

	if err := checkNoteContent(note.Content); err != nil {
		return nil, err
	}

	newNote, err := noteRepo.Save(&domain.Note{
		Profile:      note.Profile,
		Name:         note.Name,
		Content:      note.Content,
		LastModified: NowUTC(),
	})
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}
	return toNoteDTO(newNote), nil
}

func PutNote(id int, dto *NoteCreateDTO) (*NoteDTO, *APIError) {
	note, err := noteRepo.FindByID(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	if note == nil {
		return nil, ErrorNoteNotFound
	}
	dto.Name = strings.TrimSpace(dto.Name)
	dto.Profile = strings.TrimSpace(dto.Profile)

	if err := checkNoteName(dto.Name); err != nil {
		return nil, err
	}

	if err := checkNoteProfile(dto.Profile); err != nil {
		return nil, err
	}

	if err := checkNoteContent(dto.Content); err != nil {
		return nil, err
	}

	note.Profile = dto.Profile
	note.Name = dto.Name
	note.Content = dto.Content
	note.LastModified = NowUTC()
	newNote, err := noteRepo.Update(note)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}
	return toNoteDTO(newNote), nil
}

func DeleteNote(id int) *APIError {
	note, err := noteRepo.FindByID(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err)
		return ErrorInternalServer
	}

	if note == nil {
		return ErrorNoteNotFound
	}

	err = noteRepo.DeleteByID(id)
	if err != nil {
		fmt.Println(err)
		return ErrorInternalServer
	}
	return nil
}

func checkNoteName(name string) *APIError {
	length := len(name)

	if length < noteNameMinLength || length > noteNameMaxLength {
		return ErrorInvalidNoteNameRange
	}
	return nil
}

func checkNoteProfile(profile string) *APIError {
	length := len(profile)

	if length < noteProfileMinlength || length > noteProfileMaxlength {
		return ErrorInvalidNoteProfileRange
	}
	return nil
}

func checkNoteContent(content string) *APIError {
	length := len(content)

	if length > noteContentMaxLength {
		return ErrorNoteContentTooLong
	}
	return nil
}

func toNoteDTO(p *domain.Note) *NoteDTO {
	return &NoteDTO{
		ID:           p.ID,
		Profile:      p.Profile,
		Name:         p.Name,
		Content:      p.Content,
		LastModified: FormatEpoch(p.LastModified),
	}
}
