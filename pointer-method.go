package main

import "fmt"

type Man struct {
	Name string
}

// func (man Man) Married(){ // Tidak merubah data di lane 19-21
// 	man.Name = "Mr. " + man.Name
// 	fmt.Println("ini di mathod ",man.Name)
// }
func (man *Man) Married(){ // dengan pointer merubah data di lane 19-21
	man.Name = "Mr. " + man.Name
	fmt.Println("ini di mathod ",man.Name)
}

func main(){
	eko := Man{"Eko"}
	eko.Married()
	fmt.Println(eko.Name)
}