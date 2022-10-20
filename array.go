package main

import "fmt"

func main(){
	var names [3]string
	names[0] = "kurao"
	names[1] = "hikari"
	names[2] = "dimas"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var numbers = [3]int{
		90,95,80,
	}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(len(names))

}