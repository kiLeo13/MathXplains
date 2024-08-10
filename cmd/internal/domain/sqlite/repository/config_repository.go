package repository

import (
	"database/sql"
	"strconv"
)

type ConfigRepository struct {
	db *sql.DB
}

func NewConfigRepository(db *sql.DB) *ConfigRepository {
	return &ConfigRepository{db}
}

func (c *ConfigRepository) Get(key string) (string, error) {
	row := c.db.QueryRow("SELECT value FROM config WHERE key = ?", key)
	var value string

	err := row.Scan(&value)
	return value, err
}

func (c *ConfigRepository) GetInt(key string) (int, error) {
	res, err := c.Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func (c *ConfigRepository) Set(key, value string) error {
	_, err := c.db.Exec("INSERT OR REPLACE INTO config (key, value) VALUES (?, ?)", key, value)
	return err
}

// PatchInt first attempts to create a new row with 0 value,
// if the key already exists, it updates the row with the provided value.
// Negative values are allowed.
func (c *ConfigRepository) PatchInt(key string, value int) (int, error) {
	_, err := c.db.Exec(`INSERT INTO config (key, value) VALUES (?, ?)
	ON CONFLICT(key) DO UPDATE SET value = CAST(value AS INTEGER) + ?;`, key, value, value)
	if err != nil {
		return 0, err
	}
	return c.GetInt(key)
}
