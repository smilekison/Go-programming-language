package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

var events []Event = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	          VALUES (?,?,?,?,?)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare insert statement: %w", err)
	}
	defer stmt.Close()

	// Execute the insert query
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("failed to execute insert statement: %w", err)
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	// Set the ID of the event and append to events
	e.ID = id
	events = append(events, *e) // Assuming 'events' is a global slice

	// Log the result for debugging
	fmt.Println("Event saved successfully with ID:", e.ID)

	return nil // Return nil to indicate success
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, dateTime, user_id FROM events"
	rows, err := db.DB.Query(query)

	// Check if the query returned an error
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	defer rows.Close()

	var events []Event
	rowCount := 0 // Counter to track number of rows

	// Loop through the rows and scan each one into an Event struct
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event row: %w", err)
		}
		events = append(events, event)
		rowCount++
	}

	// Print row count to debug
	fmt.Println("Number of events found: ", rowCount)

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return events, nil
}
func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to scan event row: %w", err)
	}
	return &event, nil
}

func (e Event) Update() error {
	query := `UPDATE events SET name = ?, description =?, location = ?, dateTime = ? WHERE id = ?`

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare update  statement: %w", err)
	}
	fmt.Println(stmt)

	// fmt.Println("This is Name:: ", e.Name)
	// fmt.Println("This is Name:: ", e.Description)
	// fmt.Println("This is Name:: ", e.Location)
	// fmt.Println("This is Name:: ", e.DateTime)
	// fmt.Println("This is Name:: ", e.UserID)
	// Execute the insert query
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return fmt.Errorf("failed to execute insert statement: %w", err)
	}
	defer stmt.Close()
	return nil

}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	// Prepare the statement
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare delete  statement: %w", err)
	}
	fmt.Println(stmt)

	defer stmt.Close()
	// Execute the insert query
	_, err = stmt.Exec(e.ID)
	// if err != nil {
	// 	return fmt.Errorf("failed to execute delete statement: %w", err)
	// }
	return err
}
