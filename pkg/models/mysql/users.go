package mysql

import (
	"database/sql"
	"strings"

	"github.com/Soyaib10/snippetbox/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// Define a Users type which wraps a sql.DB connection pool. Needs for every table
type UserModel struct {
	DB *sql.DB
}

// Insert method to add a new record to the users table
func (m *UserModel) Insert(name, email, password string) error {
	// Create bcrypt hash of the plain text password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created)
VALUES(?, ?, ?, UTC_TIMESTAMP())`

	// Exec() method to insert the user details and hashed password into the users table.
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users.email") {
				return models.ErrDuplicateEmail
			}
		}
	}
	return err
}

// Get method to fetch details for a specific user based on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}

// Authenticate method to verify whether a user exists with the provided email address and password. This will return the relevant user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}
