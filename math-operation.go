package main

import "fmt"

func main(){
	var result = 10 + 10
	fmt.Println(result)
	var a = 10
	var b = 20

	var c = a + b
	c += result
	
	fmt.Println(c)
}