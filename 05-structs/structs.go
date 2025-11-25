package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY)")

	// ... do something awesome with that gathered data!
	u, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Print(err)
		panic("Invalid output!")
	}

	admin := user.NewAdmin("test@example.com", "test123", *u) // why u doesn't compile?

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	u.OutputUserDetails()
	u.ClearUserName()
	u.OutputUserDetails()
}

func getUserData(prompText string) string {
	fmt.Print(prompText)
	var value string
	fmt.Scanln(&value)
	return value
}
