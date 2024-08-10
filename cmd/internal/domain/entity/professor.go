package domain

type Professor struct {
	ID       int
	Name     string
	FullName string
	Nick     *string
	Known    int // This field represents a boolean column (1 = True)
}
