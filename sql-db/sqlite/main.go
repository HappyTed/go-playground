package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	// CREATE
	userID := createUser(db, "Alice", 25)
	fmt.Println("Created user with ID:", userID)

	// READ
	users := getUsers(db)
	fmt.Println("Users:")
	for _, u := range users {
		fmt.Println(u)
	}

	// UPDATE
	updateUser(db, userID, "Alice Updated", 26)

	// DELETE
	deleteUser(db, userID)
}

func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(db *sql.DB, name string, age int) int64 {
	result, err := db.Exec(
		"INSERT INTO users(name, age) VALUES (?, ?)",
		name, age,
	)
	if err != nil {
		log.Fatal(err)
	}

	id, _ := result.LastInsertId()
	return id
}

func getUsers(db *sql.DB) []User {
	query := "SELECT id, name, age FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Age)
		users = append(users, u)
	}
	return users
}

func updateUser(db *sql.DB, id int64, name string, age int) {
	query := "UPDATE users SET name = ?, age = ? WHERE id = ?"
	_, err := db.Exec(
		query,
		name, age, id,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteUser(db *sql.DB, id int64) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
}
