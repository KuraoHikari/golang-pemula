package main

import "fmt"

func main(){
	var kurao = 90
	var hikari = 80

	var kuraopoin = kurao > 80
	var hikaripoin = hikari > 80

	var kuraoHikari = kuraopoin && hikaripoin
	var kuraoHikaris = kuraopoin || hikaripoin

	fmt.Println(kuraoHikari, kuraoHikaris)
}