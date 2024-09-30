package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, _ := todo.New(todoText)

	todo.Display()
	todo.Save()
	fmt.Println("The tododododododododod !!")

	userNote, _ := note.New(title, content)
	userNote.Display()
	err := userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed. ")
		return
	}
	fmt.Println("The note has been saved successfully !!")
}

func getTodoData() string {
	return getUserInput("Todo text: ")
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
