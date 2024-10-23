package models

import (
	"errors"
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
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
	hashedPassword, err := utils.HashPassword(u.Password)

	result, _ := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println("Unable to register user. ", err)
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err

}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password from users where email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	// fmt.Println("This is boolean value for password validation::: ", passwordIsValid)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
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
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
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
