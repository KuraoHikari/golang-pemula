package main

import "fmt"

func NewMap (name string) map[string]string{
	if name == ""{
		return nil
	}else {
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	var person map[string]string
	
	if person["name"] == "" {
		fmt.Println("Data Kosong")
	}else {
		fmt.Println(person)
	}

	// dengan nil bisa tanpa mengetahui name
	var person2 map[string]string = NewMap("")
	var person3 map[string]string = NewMap("kurao")
	if person == nil {
		fmt.Println("Data Kosong")
	}else {
		fmt.Println(person2)
		fmt.Println(person3)
	}
}