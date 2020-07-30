package sqlite3

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	db, err := NewClient("./foo.db")
	if err != nil {
		t.Error(err)
	}

	db.db.Close()
}
