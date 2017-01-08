// Package user provides access to the user table in the Postgres database.
package user

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

var (
	// table is the table name.
	// postgres reserves the word 'user' so updated to app_user
	table = "app_user"
)

// Item defines the model.
type Item struct {
	ID        uint32      `db:"id"`
	FirstName string      `db:"first_name"`
	LastName  string      `db:"last_name"`
	Email     string      `db:"email"`
	Password  string      `db:"password"`
	StatusID  uint8       `db:"status_id"`
	CreatedAt pq.NullTime `db:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at"`
}

// Connection is an interface for making queries.
type Connection interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

// ByEmail gets user information from email.
func ByEmail(db Connection, email string) (Item, bool, error) {
	result := Item{}

	err := db.Get(&result, fmt.Sprintf(`
		SELECT id, password, status_id, first_name
		FROM %v
		WHERE email = $1
		AND deleted_at IS NULL
		LIMIT 1
		`, table),
		email)

	return result, err == sql.ErrNoRows, err
}

// Create creates user.
func Create(db Connection, firstName, lastName, email, password string) (sql.Result, error) {
	result, err := db.Exec(fmt.Sprintf(`
		INSERT INTO %v
		(first_name, last_name, email, password)
		VALUES
		($1,$2,$3,$4)
		`, table),
		firstName, lastName, email, password)
	return result, err
}
