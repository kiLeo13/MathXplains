package repository

import (
	"MathXplains/internal/domain/entity"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Save(id, name string, timestamp int64) error {
	_, err := u.db.Exec(`INSERT INTO users (id, name, admin, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)`, id, name, 0, timestamp, timestamp)
	return err
}

func (u *UserRepository) IsAdmin(userId string) (bool, error) {
	res := u.db.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE id = ? AND admin = 1)", userId)

	var value int
	err := res.Scan(&value)
	if err != nil {
		return false, err
	}

	return value == 1, nil
}

func (u *UserRepository) Verify(userId string) error {
	_, err := u.db.Exec("UPDATE users SET email_verified = 1 WHERE id = ?", userId)
	return err
}

func (u *UserRepository) FindAll() ([]*domain.User, error) {
	res, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	return serializeUsers(res)
}

func (u *UserRepository) FindById(id string) (*domain.User, bool, error) {
	res := u.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user, err := serializeUser(res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		} else {
			return nil, false, err
		}
	}
	return user, true, nil
}

func serializeUsers(rows *sql.Rows) ([]*domain.User, error) {
	var users []*domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Admin,
			&user.EmailVerified,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func serializeUser(row *sql.Row) (*domain.User, error) {
	var user domain.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Admin,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return &user, err
}
