package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int)(int , error){
	if pembagi == 0 {
		return 0, errors.New("Pembagi ga boleh nol")

	} else {
		res := nilai / pembagi
		return res , nil
	}
}

func main(){
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil , err := Pembagi(100 , 0)

	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}