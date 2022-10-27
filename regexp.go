package main

import ("fmt"
		"regexp")

func main(){
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	// fmt.Println("Hello World")
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("elo"))
	fmt.Println(regex.MatchString("eFo"))

	search := regex.FindAllString("eko eka edo eto eyo eki", -1)
	fmt.Println(search)
}