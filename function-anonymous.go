package main

import "fmt"

type Blacklist func(string) bool

func regisUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	regisUser("admin" , blacklist)
	regisUser("kurao", blacklist)
	regisUser("kurao", func(s string) bool {
		return s == "kurao"
	})
}