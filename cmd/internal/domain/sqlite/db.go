package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type RowScanner interface {
	Scan(dest ...any) error
}

func Init() (*sql.DB, error) {
	conn, err := newConnection()
	if err != nil {
		return nil, err
	}
	tables := getTables()

	for _, table := range tables {
		_, err := conn.Exec("CREATE TABLE IF NOT EXISTS " + table)
		if err != nil {
			return conn, err
		}
	}
	return conn, nil
}

func newConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", os.Getenv("SQLITE_URL"))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func getTables() []string {
	return []string{
		`users (
    	  id TEXT PRIMARY KEY,
    	  name TEXT NOT NULL,
    	  admin INTEGER NOT NULL DEFAULT 0,
    	  created_at BIGINT NOT NULL,
    	  updated_at BIGINT NOT NULL
		)`,

		`professors (
    	  id INTEGER PRIMARY KEY AUTOINCREMENT,
    	  name TEXT NOT NULL,
    	  full_name TEXT NOT NULL,
    	  nick TEXT,
    	  known INTEGER NOT NULL DEFAULT 0
		)`,

		`subjects (
          id INTEGER PRIMARY KEY AUTOINCREMENT,
          professor_id INTEGER,
          name TEXT NOT NULL,
          full_name TEXT NOT NULL,
		  available INTEGER NOT NULL DEFAULT 0,
          FOREIGN KEY(professor_id) REFERENCES professors(id)
        )`,

		`appointments (
    	  id INTEGER PRIMARY KEY AUTOINCREMENT,
    	  topic TEXT NOT NULL,
    	  description TEXT NOT NULL,
    	  user_id TEXT NOT NULL,
    	  subject_id BIGINT NOT NULL,
    	  professor_id BIGINT,
		  rejected INTEGER NOT NULL DEFAULT 0,
		  scheduled_at BIGINT NOT NULL,
    	  created_at BIGINT NOT NULL,
		  updated_at BIGINT NOT NULL,
		  active INTEGER NOT NULL DEFAULT 1,
    	  FOREIGN KEY(user_id) REFERENCES users(id),
    	  FOREIGN KEY(subject_id) REFERENCES subjects(id),
    	  FOREIGN KEY(professor_id) REFERENCES professors(id)
		)`,

		`notes (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  profile TEXT NOT NULL,
		  name TEXT NOT NULL,
		  content TEXT NOT NULL,
		  created_at BIGINT NOT NULL,
		  last_modified BIGINT NOT NULL,
		  active INTEGER NOT NULL DEFAULT 1
		)`,
	}
}
