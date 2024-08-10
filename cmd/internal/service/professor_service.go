package service

import (
	"fmt"
)

type ProfessorDTO struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	FullName string  `json:"full_name"`
	Nick     *string `json:"nick"`
	Known    bool    `json:"known"`
}

func GetProfessors(knownOnly bool) []*ProfessorDTO {
	professors, err := professorRepo.FindAll(knownOnly)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var professorList []*ProfessorDTO
	for _, p := range professors {
		professorList = append(professorList, &ProfessorDTO{
			ID:       p.ID,
			Name:     p.Name,
			FullName: p.FullName,
			Nick:     p.Nick,
			Known:    p.Known == 1,
		})
	}
	return professorList
}
