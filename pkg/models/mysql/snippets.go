package mysql

import (
	"database/sql"

	"github.com/Soyaib10/snippetbox/pkg/models"
)

// Define a SnippetModel type which wraps a sql.DB connection pool. Needs for every table
type SnippetModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets(title, content, created, expires)
			VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Single record SQL
	row := m.DB.QueryRow(stmt, id) // Returns a single row
	s := &models.Snippet{}         // Initialize a pointer to a new zeroed Snippet struct.

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord // var ErrNoRecord = errors.New("models: no matching record found")
	} else if err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippet object.
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires
			 FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	rows, err := m.DB.Query(stmt) // returns multiple rows
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure the sql.Rows resultset is always properly closed before the Latest() method returns

	snippets := []*models.Snippet{} // Initialize an empty slice to hold the models.Snippets objects because it retruns multiple rows
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s) 
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
