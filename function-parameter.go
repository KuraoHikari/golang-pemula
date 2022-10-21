package main

import "fmt"

func main(){
	sayHello("Kurao", "WIBU")
}

func sayHello(name string , title string){
	fmt.Println("Hello ", name , " si ", title)
}