package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "eko"
	middleName = "kurao"
	lastName = "hikari"
	return
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}