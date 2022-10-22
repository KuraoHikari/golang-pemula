package main

import "fmt"

func main(){
	counter := 0 
	increment := func(){
		counter++ 
	}
	increment()
	increment()
	fmt.Println(counter)

	name := "kurao"

	changger := func(){
		name:= "Budi"
		fmt.Println(name)
	}
	fmt.Println(name)
	changger()
}