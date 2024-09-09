package mysql

import (
	"database/sql"

	"github.com/Soyaib10/snippetbox/pkg/models"
)

// Define a Users type which wraps a sql.DB connection pool. Needs for every table
type Users struct {
	DB *sql.DB
}

// Insert method to add a new record to the users table
func (m *Users) Insert(name, email, password string) error {
	return nil
}

// Get method to fetch details for a specific user based on their user ID.
func (m *Users) Get(id int) (*models.User, error) {
	return nil, nil
}

// Authenticate method to verify whether a user exists with the provided email address and password. This will return the relevant user ID if they do.
func (m *Users) Authenticate(email, password string) (int, error) {
	return 0, nil
}
