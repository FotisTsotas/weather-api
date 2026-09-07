package models

type User struct {
	ID       int64
	Username string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}
