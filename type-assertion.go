package main

import "fmt"

func random() interface{}{
	return 10
	//return true //panic jika lane 13-14 di jalankan
}

func main(){
	var result interface{} = random()

	// var resultString string = result.(string)
	// fmt.Println(resultString)

	//lebih aman jika memakai switch untuk mengetahui tipe data lalu di eksekusi
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "string")
	case int:
		fmt.Println("Value", value, "int")
	default:
		fmt.Println("Uknown")
	}
	
}