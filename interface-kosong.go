package main

import "fmt"

func Ups(i int) interface{}{
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main(){
	// var data int = Ups(1) <== tidak bisa 
	// fmt.Println(data)
	var data interface{} = Ups(1)
	fmt.Println(data)
}