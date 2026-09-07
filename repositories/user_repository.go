package repositories

import (
	"database/sql"
	"weather-api/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = ?"
	row := r.db.QueryRow(query, email)

	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
