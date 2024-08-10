package domain

type Subject struct {
	ID          int
	ProfessorID *int
	Name        string
	FullName    string
	Available   int // This field represents a boolean column (1 = True)
}
