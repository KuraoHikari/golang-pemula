package main

import "fmt"

func main(){
	var (
		nilai32 int32 = 100000
		nilai64 int64 = int64(nilai32)
		nilai8 int8 = int8(nilai64)
		nilai16 int16 = int16(nilai32)
	)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)
	fmt.Println(nilai16)

	var (
		name = "Kurao"
		e byte = name[0]
		eString string = string(e)
		aString = string(name[2])
	)
	
	fmt.Println(eString)
	fmt.Println(aString)
}