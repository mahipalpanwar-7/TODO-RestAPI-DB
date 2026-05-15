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

	query := `INSERT INTO users(username,email,password)VALUES(?,?,?)`

	config.DB.Exec(
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

