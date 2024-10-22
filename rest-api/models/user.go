package models

import (
	"fmt"

	"example.com/rest-api/db"
)

type User struct {
	Id       int64
	Email    string
	Password string
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Unable to register user. ", err)
	}
	defer stmt.Close()

	result, _ := stmt.Exec(u.Email, u.Password)
	if err != nil {
		fmt.Println("Unable to register user. ", err)
	}
	userId, err := result.LastInsertId()
	u.Id = userId
	return err

}
func GetAllUsers() ([]User, error) {
	query := "SELECT id, email, password FROM users"
	rows, err := db.DB.Query(query)

	// Check if the query returned an error
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	var users []User
	rowCount := 0 // Counter to track number of rows

	// Loop through the rows and scan each one into an Event struct
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event row: %w", err)
		}
		users = append(users, user)
		rowCount++
	}

	// Print row count to debug
	fmt.Println("Number of events found: ", rowCount)

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return users, nil
}
