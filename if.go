package main

import "fmt"

func main(){
	name:= "kurao"
	if name == "krao" {
		fmt.Println("masuk")
	} else if name == "kurao" {
		fmt.Println("masuk ke tengah")
	} else {
		fmt.Println("ga masuk")
	}

	if length := len(name); length > 5 {
		fmt.Println("masuk dengan short statment")
	} else {
		fmt.Println("ga masuk dengan short statment")
	}
	if len(name) > 3 {
		fmt.Println("masuk dengan short statment")
	} else {
		fmt.Println("ga masuk dengan short statment")
	}
}