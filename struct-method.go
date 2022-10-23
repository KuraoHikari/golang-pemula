package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string){ //struct method
	fmt.Println("Hello", name, "My Name is", customer.Name)
}
func (a Customer) sayYooo(name string){ //struct method
	fmt.Println("Say yooo", name, "Chika dance", a.Name)
}

func main() {
	var kurao Customer
	kurao.Name = "kurao hikari"
	kurao.Address = "Indonesia"
	kurao.Age = 18

	kurao.sayHi("moe moe kyun")
	kurao.sayYooo("moe moe kyun")
	// fmt.Println(kurao)
	// fmt.Println(kurao.Age)

	// var hikari = Customer{
	// 	Name: "Joko",
	// 	Address: "Indonesia",
	// 	Age: 99,
	// }
	// fmt.Println(hikari)
	// fmt.Println(hikari.Age)
}