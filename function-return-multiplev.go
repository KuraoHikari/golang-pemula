package main

import "fmt"

func main() {
	return1, return2 := sayHello("kurao", "wibu")
	return3, _ := sayHello("kurao", "wibu")
	fmt.Println(return1 +" == " + return2 ,return3)
}

func sayHello(name string, title string) (string, string) {
	if name == "" {
		return "nani", title
	} else {
		return "Hello " + name, title
	}
}