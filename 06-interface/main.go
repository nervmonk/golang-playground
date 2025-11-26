package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
	"example.com/structs-practice/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething(true)
	printSomething("Hello World")
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(userNote)
}

func printSomething(value any) {
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("It is integer: ", typedVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println(value)
	// default:
	// 	fmt.Println("Type not known")
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
