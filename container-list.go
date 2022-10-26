package main

import ("fmt"
		"container/list")

func main(){
	data := list.New()

	data.PushBack("Kurao")
	data.PushBack("Hikari")
	data.PushBack("Annnjay")
	data.PushFront("Wibu")

	// data.Front().Next().Next()
	// data.Back().Prev().Prev()

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil ; element = element.Next(){
		fmt.Println(element.Value)
	}

	for element := data.Back(); element != nil ; element = element.Prev(){
		fmt.Println(element.Value)
	}
}