package mysql

import (
	"database/sql"

	"swarmtree.com/pegion/pkg/models"
)

type PegionModel struct {
	DB *sql.DB
}

// Insert new pegion into the database
func (m *PegionModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO pegions (title, content, created, expires)
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

// Return a specified pegion based on its id
func (m *PegionModel) Get(id int) (*models.Pegion, error) {

	stmt := `SELECT id, title, content, created, expires FROM pegions
			WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Pegion{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// Return 10 latest pegions
func (m *PegionModel) Latest() ([]*models.Pegion, error) {
	return nil, nil
}
