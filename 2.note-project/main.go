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

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	save() error
// 	Display()
// }

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	userNote, err := note.New(title, content)
	// todo.Display()
	outputData(todo)
	if err != nil {
		return
	}
	// userNote.Display()
	outputData(userNote)
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

func outputData(data outputtable) {
	data.Display()
	saveData(data)
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
