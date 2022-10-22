package main

import "fmt"

func sumAll(name string, numbers ...int)int {
	total :=0
	for _, value := range numbers {
		total += value
	}
	fmt.Println(name)
	return total
}

func main(){
	total :=sumAll("kurao",1,2,3,4,5)
	fmt.Println(total)

	slice := []int{10,20,30,40,50}
	total = sumAll("hikari", slice...)
	fmt.Println(total)
}