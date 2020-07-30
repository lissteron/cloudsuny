package sqlite3

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3" // driver
)

type Client struct {
	db *sqlx.DB
}

func NewClient(path string) (*Client, error) {
	db, err := sqlx.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	client := &Client{db: db}

	if err := client.initDB(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) DB() *sqlx.DB {
	return c.db
}

func (c *Client) initDB() error {
	query := `
	CREATE TABLE IF NOT EXISTS "user" (
		id UUID NOT NULL PRIMARY KEY,
		username TEXT NOT NULL,
		avatar TEXT NOT NULL DEFAULT '',
		created_at TIMESTAMP NOT NULL
	);
	`

	if _, err := c.db.Exec(query); err != nil {
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS "badge" (
		id UUID NOT NULL PRIMARY KEY,
		user_id UUID NOT NULL,
		type TEXT NOT NULL,
		x NUMERIC NOT NULL DEFAULT 0,
		y NUMERIC NOT NULL DEFAULT 0,
		created_at TIMESTAMP NOT NULL,
		CONSTRAINT fk_badge_ref_user FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE
	);`

	if _, err := c.db.Exec(query); err != nil {
		return err
	}

	query = `CREATE UNIQUE INDEX IF NOT EXISTS user_username_key ON "user"(username);`

	if _, err := c.db.Exec(query); err != nil {
		return err
	}

	return nil
}
