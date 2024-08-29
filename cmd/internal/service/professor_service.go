package service

import (
	domain "MathXplains/internal/domain/entity"
	"fmt"
)

type ProfessorDTO struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	FullName string  `json:"full_name"`
	Nick     *string `json:"nick"`
	Known    bool    `json:"known"`
}

func GetProfessors(knownOnly bool) ([]*ProfessorDTO, *APIError) {
	professors, err := professorRepo.FindAll(knownOnly)
	if err != nil {
		fmt.Println(err)
		return nil, ErrorInternalServer
	}

	var professorList []*ProfessorDTO
	for _, p := range professors {
		professorList = append(professorList, toProfessorDTO(p))
	}
	return professorList, nil
}

func toProfessorDTO(p *domain.Professor) *ProfessorDTO {
	return &ProfessorDTO{
		ID:       p.ID,
		Name:     p.Name,
		FullName: p.FullName,
		Nick:     p.Nick,
		Known:    p.Known == 1,
	}
}
