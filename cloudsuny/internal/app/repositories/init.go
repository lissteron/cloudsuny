package repositories

import (
	"github.com/lissteron/simplerr"
	"github.com/loghole/database"
)

func Migrate(db *database.DB) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS "user" (
			id UUID NOT NULL PRIMARY KEY,
			username TEXT NOT NULL,
			avatar TEXT NOT NULL DEFAULT '',
			created_at TIMESTAMP NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS "badge" (
			id UUID NOT NULL PRIMARY KEY,
			user_id UUID NOT NULL,
			type TEXT NOT NULL,
			x NUMERIC NOT NULL DEFAULT 0,
			y NUMERIC NOT NULL DEFAULT 0,
			created_at TIMESTAMP NOT NULL,
			CONSTRAINT fk_badge_ref_user FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE RESTRICT ON UPDATE CASCADE
		);`,
		`CREATE UNIQUE INDEX IF NOT EXISTS user_username_key ON "user"(username);`,
	}

	for _, query := range migrations {
		if _, err := db.Exec(query); err != nil {
			return simplerr.Wrapf(err, "query '%s' failed", query)
		}
	}

	return nil
}
