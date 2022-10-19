package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool

	var kurao NoKTP = "Hikari"
	var status Married = true
	
	fmt.Println(kurao)
	fmt.Println(status)
}