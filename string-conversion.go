package main

import ("fmt"
		"strconv")

func main(){
	boolean , err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	}else {
		fmt.Println(err.Error())
	}

	number , err := strconv.ParseInt("1000000s", 10, 64) // selain ParseInt ada Atoi dan Itoa
	if err == nil {
		fmt.Println(number)
	}else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000,8)
	fmt.Println(value)

	valueAtoi,_ := strconv.Atoi("100009")
	fmt.Println(valueAtoi)
	valueItoa := strconv.Itoa(100555)
	fmt.Println(valueItoa)
}