package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/config"
	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/models"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := `SELECT id,title,is_completed FROM todos`

	rows, err := config.DB.Query(query)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {

		var todo models.Todo

		rows.Scan(
			&todo.Id,
			&todo.Title,
			&todo.IsCompleted,
		)

		todos = append(todos, todo)
	}

	json.NewEncoder(w).Encode(todos)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		json.NewEncoder(w).Encode("Invalid request")
		return
	}

	query := `INSERT INTO todos(title,is_completed)
	VALUES(?,?)`

	_, err = config.DB.Exec(query, todo.Title, false)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode("Todo added successfully")

}

func CompleteTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid JSON or is_completed must be true/false")
		return
	}

	query := `UPDATE todos SET is_completed=? WHERE ID=?`

	_, err = config.DB.Exec(query, todo.IsCompleted, id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	} else {
		json.NewEncoder(w).Encode("todo completed")
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode("invalid id")
		return
	}

	query := `DELETE FROM todos Where id=?`

	_, err = config.DB.Exec(query, id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode("Todo deleted successfully")

}
