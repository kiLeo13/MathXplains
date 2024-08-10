package service

import (
	"fmt"
)

type SubjectDTO struct {
	ID          int    `json:"id"`
	ProfessorID *int   `json:"professor_id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Available   bool   `json:"available"`
}

func GetSubjects(available bool) []*SubjectDTO {
	subjects, err := subjectRepo.FindAll(available)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var subjectList []*SubjectDTO
	for _, s := range subjects {
		subjectList = append(subjectList, &SubjectDTO{
			ID:          s.ID,
			ProfessorID: s.ProfessorID,
			Name:        s.Name,
			FullName:    s.FullName,
			Available:   s.Available == 1,
		})
	}
	return subjectList
}
