package service

import (
	domain "MathXplains/internal/domain/entity"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	noteNameMinLength = 1
	noteNameMaxLength = 30

	noteProfileMinlength = 3
	noteProfileMaxlength = 50

	noteContentMaxLength = 500_000
)

type NoteDTO struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	CreatedAt    string `json:"created_at"`
	LastModified string `json:"last_modified"`
}

type NoteCreateDTO struct {
	Profile string `json:"profile"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// NotePreviewDTO is the entity used for the responses of GetNotesSummary,
// returing the fewest amount of data as possible,
// so the client is aware of all the notes they have created.
//
// A payload containing all the data (mostly the content) might overwhelm the
// client and use unnecessary memory and bandwidth.
//
// This is the same principle of listing all the files in a file explorer,
// you don't have to fully open them.
type NotePreviewDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (n *NoteDTO) toPreview() *NotePreviewDTO {
	return &NotePreviewDTO{
		ID:   n.ID,
		Name: n.Name,
	}
}

func GetNote(profile string, id int) (*NoteDTO, *APIError) {
	hashedProfile := HashProfile(profile)
	note, err := noteRepo.FindByProfileAndID(hashedProfile, id)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	if note == nil {
		return nil, ErrorNoteNotFound
	}
	return toNoteDTO(note), nil
}

func GetNotesSummary(profile string) ([]*NotePreviewDTO, *APIError) {
	hashedProfile := HashProfile(profile)
	notes, err := noteRepo.FindAllByProfile(hashedProfile)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	noteList := make([]*NotePreviewDTO, len(notes))
	for i, n := range notes {
		noteList[i] = toNoteDTO(n).toPreview()
	}
	return noteList, nil
}

func CreateNote(note *NoteCreateDTO) (*NoteDTO, *APIError) {
	note.Name = strings.TrimSpace(note.Name)
	note.Profile = strings.TrimSpace(note.Profile)
	hashedProfile := HashProfile(note.Profile)
	found, err := noteRepo.ExistsByProfileAndName(hashedProfile, note.Name)
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

	now := NowUTC()
	newNote, err := noteRepo.Save(&domain.Note{
		Profile:      hashedProfile,
		Name:         note.Name,
		Content:      note.Content,
		CreatedAt:    now,
		LastModified: now,
	})
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}
	return toNoteDTO(newNote), nil
}

func PutNote(profile string, id int, dto *NoteCreateDTO) (*NoteDTO, *APIError) {
	if strings.TrimSpace(profile) == "" {
		return nil, ErrorProfileNotProvided
	}
	hashedProfile := HashProfile(profile)
	note, err := noteRepo.FindByProfileAndID(hashedProfile, id)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	if note == nil {
		return nil, ErrorNoteNotFound
	}
	dto.Name = strings.TrimSpace(dto.Name)

	if err := checkNoteName(dto.Name); err != nil {
		return nil, err
	}

	if err := checkNoteContent(dto.Content); err != nil {
		return nil, err
	}

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

func DeleteNote(profile string, id int) *APIError {
	if strings.TrimSpace(profile) == "" {
		return ErrorProfileNotProvided
	}
	hashedProfile := HashProfile(profile)
	note, err := noteRepo.FindByProfileAndID(hashedProfile, id)
	if err != nil {
		fmt.Println(err)
		return ErrorInternalServer
	}

	if note == nil {
		return ErrorNoteNotFound
	}

	err = noteRepo.DeleteByID(note.ID)
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
		Name:         p.Name,
		Content:      p.Content,
		CreatedAt:    FormatEpoch(p.CreatedAt),
		LastModified: FormatEpoch(p.LastModified),
	}
}

func HashProfile(profile string) string {
	hasher := sha512.New()

	hasher.Write([]byte(profile))
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
