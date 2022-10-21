package main

import "fmt"

func main() {
	res := sayHello("", "Hikari")
	fmt.Println(res)
	fmt.Println(sayHello("kurao", "Hikari"))
}

func sayHello(name string, title string) string {
	if name == "" {
		return "nani"
	} else {
		return "Hello " + name
	}
}