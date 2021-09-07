package controllers

import (
	"log"

	"github.com/Henry-Asante/auth_service/models"
)

// AuthContoller typing for other files
type AuthContoller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (u AuthContoller) Signup(db, user models.User) models.User {
	// stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	// err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)

	// logFatal(err)

	// user.Password = ""
	return user
}

func (u AuthContoller) Login(db, user models.User) (models.User, error) {
	// row := db.QueryRow("select * from users where email=$1", user.Email)
	// err := row.Scan(&user.ID, &user.Email, &user.Password)

	// if err != nil {
	// 	return user, err
	// }

	return user, nil
}
