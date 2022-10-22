package main

import "fmt"

type filterType func(string) string
func sayHelloWithFilter(name string, filter filterType){
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}
func spamFilter(name string) string {
	if name == "anjing"{
		return "..."
	} else {
		return name
	}
}

func main(){
	sayHelloWithFilter("kurao", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}