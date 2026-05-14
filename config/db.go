package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := ""

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("database connected successfully")

	DB = db

	CreateDatabase()
	createUserTable()
	createTodosTable()

}

func CreateDatabase() {
	query := `CREATE TABLE IF NOT EXISTS todoapp`
	_, err := DB.Exec(query)

	if err != nil {
		panic(err)
	}
	fmt.Println("database created successfully")
}

func createUserTable() {
	query := `CREATE TABLE IF NOT EXISTS user(
	id INT PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(100),
	email VARCHAR(100) UNIQUE,
	password VARCHAR(255)
	)`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
	fmt.Println("user table created successfully")
}

func createTodosTable() {
	query := `CREATE TABLE IF NOT EXISTS todos(
	id INT PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(255),
	is_completed BOOLEAN DEFAULT FALSE
	)`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
   fmt.Println("todo table created successfully")
}

