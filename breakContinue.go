package main

import "fmt"

func main(){
	for i := 0 ; i<10 ; i++ {
		fmt.Println("loop ", i)
		if i == 2 {
			break
		}
		fmt.Println("loop ", i)
	}
	for i := 0 ; i < 10 ; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("loop ", i)
	}
}