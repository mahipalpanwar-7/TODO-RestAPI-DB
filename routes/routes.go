package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/controllers"
	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// public routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	//protected routes
	router.Handle("/todos", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetTodos))).Methods("GET")
	router.Handle("/todo", middleware.AuthMiddleware(http.HandlerFunc(controllers.AddTodo))).Methods("POST")
	router.Handle("/todo/{id}", middleware.AuthMiddleware(http.HandlerFunc(controllers.CompleteTodo))).Methods("PUT")
	router.Handle("/todos/{id}", middleware.AuthMiddleware(http.HandlerFunc(controllers.DeleteTodo))).Methods("DELETE")

	return router
}
