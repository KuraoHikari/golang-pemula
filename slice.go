package main

import "fmt"

func main(){
	var months = [...]string {
		"januari",
		"february",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "change"
	// slice1[2]= "yei"
	// fmt.Println(slice1)
	// fmt.Println(months)
	fmt.Println("===========================")
	var slice2 = months[8:12]
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	var slice3 = append(slice2,"woi")
	var slice4 = append(slice2,"aaaaaaaaaaaaaaa")
	fmt.Println(months)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println("===========================")

	newSlice := make([]string, 3,5)
	newSlice[0] = "kurao"
	newSlice[1] = "hikari"
	newSlice[2] = "hikari"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice),cap(newSlice))
	copy(copySlice,newSlice)
	fmt.Println(copySlice)
	var slice5 = append(newSlice,"woi")
	fmt.Println(slice5)
	fmt.Println(newSlice)
}