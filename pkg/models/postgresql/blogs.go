package postgresql

import (
	"database/sql"

	"mavila_frosado.net/test1/pkg/models"
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

func (m *BlogModel) Read() ([]*models.Blog, error) {
	query := `
	SELECT * FROM blogs`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}

	//clean before we leave Read()
	defer rows.Close()

	blogs := []*models.Blog{}

	for rows.Next() {
		blog := &models.Blog{}
		err = rows.Scan(&blog.ID, &blog.FirstName, &blog.LastName, &blog.Email, &blog.Subject, &blog.Message, &blog.Date_Created)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return blogs, nil

}
