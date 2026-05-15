package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/config"
	"github.com/mahipalpanwar-7/TODO-RestAPI-DB/routes"
)

func main() {
	config.ConnectDB()
	router := routes.Router()

	fmt.Println("server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
