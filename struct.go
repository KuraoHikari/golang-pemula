package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var kurao Customer
	kurao.Name = "kurao hikari"
	kurao.Address = "Indonesia"
	kurao.Age = 18

	fmt.Println(kurao)
	fmt.Println(kurao.Age)

	var hikari = Customer{
		Name: "Joko",
		Address: "Indonesia",
		Age: 99,
	}
	fmt.Println(hikari)
	fmt.Println(hikari.Age)
}