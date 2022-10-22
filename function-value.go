package main

import "fmt"

func getHello(name string)string{
	return "Hello ges " + name
}

func main(){
	sayHello := getHello

	fmt.Println(sayHello("kurao"))
}