package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	todo.Display()
	err = saveData(todo)
	// err := todo.Save()
	if err != nil {

		return
	}
	userNote, err := note.New(title, content)
	if err != nil {
		return
	}
	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed. ")
		return err
	}
	fmt.Println("The note has been saved successfully !!")
	return err
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Get content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n')
	if error != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	// var value string
	// fmt.Scanln(&value)
	return text
}
