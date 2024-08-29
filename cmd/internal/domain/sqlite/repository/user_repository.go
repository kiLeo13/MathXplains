package repository

import (
	"MathXplains/internal/domain/entity"
	"MathXplains/internal/domain/sqlite"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) Save(user *domain.User) error {
	_, err := u.db.Exec(`INSERT INTO users (id, name, admin, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)`, user.ID, user.Name, user.Admin, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
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

func (u *UserRepository) FindAll() ([]*domain.User, error) {
	res, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	return deserializeUsers(res)
}

// FindById searches the database using the indexed id column.
//
// Returns a nil pointer and a nil error if no users are found.
func (u *UserRepository) FindById(id string) (*domain.User, error) {
	res := u.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	user, err := deserializeUser(res)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return user, nil
}

func (u *UserRepository) DeleteByID(id string) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func deserializeUsers(rows *sql.Rows) ([]*domain.User, error) {
	var users []*domain.User

	for rows.Next() {
		user, err := deserializeUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func deserializeUser(row sqlite.RowScanner) (*domain.User, error) {
	var user domain.User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Admin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
