package main

import (
	"fmt"
)

func main(){
	counter := 1

	for counter <= 10 {
		fmt.Println("Loop ke ", counter)
		counter++
	}

	for number :=1 ; number <=10 ; number++{
		fmt.Println("Loop Number ke ", number)
	}

	slice := []string{"Eko","kurniawam","kurao","anajay"}
	for i := 0 ; i < len(slice) ; i++{
		fmt.Println("Loop Number ke ", slice[i])
	}

	names := []string{"kurao", "hikari", "hello"}

	for _, value := range names {
		fmt.Println("name" , "=",value)
	}
	for i, value := range names {
		fmt.Println("name", i , "=",value)
	}

	person := make(map[string]string)
	person["name"] = "Kurao"
	person["title"] = "Programer"

	for key, value := range person {
		fmt.Println(key, "=" , value)
	}
}