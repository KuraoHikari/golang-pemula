package main

import "fmt"

func main(){
	var name string
	name = "Kurao"
	fmt.Println(name)
	name = "hikari"
	fmt.Println(name)

	
	var string_name = "Kurao"
	fmt.Println(string_name)
	string_name = "hikari"
	fmt.Println(string_name)

	var age = 38 //tergantubng sistem operasi bisa menjadi int32 atau int64
	fmt.Println(age)
	var number int8 = 32
	fmt.Println(number)

	no_var := "Novar"
	fmt.Println(no_var)
	no_var = "re Assign Novar"
	fmt.Println(no_var)

	var (
		firstName = "kurao"
		lastName = "hikari"
		ageNumber int8 = 20
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(ageNumber)
}