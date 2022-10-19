package main

import "fmt"

func main(){
	const (
		firstName string = "Kurao"
		lastName  = "hikari"
		value = 1000
	)

	// firstName = "oaakwokwowk" //tidak bisa di re assign untuk const

	fmt.Println(firstName)
	fmt.Println(value)
}