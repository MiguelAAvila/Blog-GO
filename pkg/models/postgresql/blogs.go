package postgresql

import (
	"database/sql"
)

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(firstname, lastname, email, subject, message string) (int, error) {
	var id int
	//Write to the database
	query := `
		INSERT INTO blogs(first_name, last_name, email, subject, message)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := m.DB.QueryRow(query, firstname, lastname, email, subject, message).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil

}
