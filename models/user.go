package models

import (
	"errors"

	"arsh.com/rest-api/db"
	"arsh.com/rest-api/routes/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()

	u.ID = userId
	return err

}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassvord string
	err := row.Scan(&retrivedPassvord)

	if err != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassvord)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
