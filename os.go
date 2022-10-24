package main
import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])

	name, err := os.Hostname()

	if err == nil {
		fmt.Println("Hostname :", name)
	}else {
		fmt.Println("err", err.Error())
	}

	username := os.Getenv("USSR_USERNAME")
	password := os.Getenv("USSR_PASSWORD")

	fmt.Println("USERNAME :", username)
	fmt.Println("PASSWORD :", password)
}