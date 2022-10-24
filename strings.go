package main

import ("fmt"
		"strings")

func main(){
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Contains("Eko Kurniawan", "Ekos"))
	fmt.Println(strings.Split("Eko Kurniawan Kanedy", " "))
	fmt.Println(strings.Split("Eko Kurniawan Kanedy", "K"))
	fmt.Println(strings.ToLower("Eko Kurniawan Kanedy"))
	fmt.Println(strings.ToUpper("Eko Kurniawan Kanedy"))
	fmt.Println(strings.ToTitle("eko kurniawan kanedy"))
	fmt.Println(strings.Trim("aaaaEko kurniawanaaaa", "a"))
	fmt.Println(strings.Trim("   Eko kurniawan     ", " "))
	fmt.Println(strings.Trim("a   Eko kurniawan     a", " "))
	fmt.Println(strings.ReplaceAll("Eko Eko Kurao", "Eko", "Bedi"))
}