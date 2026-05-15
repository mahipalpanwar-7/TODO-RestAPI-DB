package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/config"
	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/models"
	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		json.NewEncoder(w).Encode("Error while hashing password")
		return
	}

	query := `INSERT INTO user(username,email,password)VALUES(?,?,?)`

	_, err = config.DB.Exec(
		query,
		user.Username,
		user.Email,
		hashedPassword,
	)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("User Registered Successfully")
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	query := `
	SELECT id,username,email,password
	FROM user
	WHERE email=?
	`

	row := config.DB.QueryRow(query, user.Email)

	var dbUser models.User

	err := row.Scan(
		&dbUser.Id,
		&dbUser.Username,
		&dbUser.Email,
		&dbUser.Password,
	)

	if err != nil {
		json.NewEncoder(w).Encode("Invalid Email")
		return
	}

	checkPassword := utils.CheckPasswordHash(
		user.Password,
		dbUser.Password,
	)

	if !checkPassword {
		json.NewEncoder(w).Encode("Invalid Password")
		return
	}

	token,err := utils.GenerateJWT(dbUser.Email)
	if err!=nil{
		json.NewEncoder(w).Encode("error while generating token")
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}
