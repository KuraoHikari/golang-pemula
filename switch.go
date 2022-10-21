package main

import "fmt"

func main(){
	name := "Kurao"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Kurao":
		fmt.Println("Hello Kurao")
	default:
		fmt.Println("anjay")
	}

	switch length := len(name); length > 5{
	case true : 
		fmt.Println("length lebih dari 5")
	case false : 
		fmt.Println("length kurang dari 5")
	}
	
	length := len(name)
	switch {
	case length > 5 : 
		fmt.Println("length lebih dari 5")
	case length <=5 : 
		fmt.Println("length kurang dari 5 di bawah")
	}
}