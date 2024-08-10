package domain

type User struct {
	ID            string
	Name          string
	Admin         bool
	EmailVerified bool
	VerifiedAt    *int64
	CreatedAt     int64
	UpdatedAt     int64
}
