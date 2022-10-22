package main

import "fmt"

func endApp(){
	message := recover()//recover di simpan di defer function
	if message != nil {
		fmt.Println("error dengan MSG" , message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp (error bool){
	defer endApp()
	if error {
		panic("APP ERROR")
	}
	
	fmt.Println("aplikasi jalan")
}
func main(){
	runApp(true)
	fmt.Println("kurao")
}