package main

import "fmt"

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApp(value int){
	defer logging()
	fmt.Println("Run APP")
	res := 10/ value
	fmt.Println("res", res)
	// logging()
}
func main(){
	runApp(10)
}