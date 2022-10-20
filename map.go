package main

import (
	"fmt"
)

func main(){
	// var username map[string]string = map[string]string{
	// 	"firstName" : "Kurao",
	// 	"lastname" : "HIkari",
	// }
	username := map[string]string{
		"firstName" : "Kurao",
		"lastname" : "HIkari",
	}
	username["password"] = "aaaa"

	fmt.Println(username)
	fmt.Println(len(username))
	fmt.Println(username["firstName"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Golang"
	book["author"] = "kurao"
	book["cover"] = "red"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "cover")

	fmt.Println(book)
	fmt.Println(len(book))
}