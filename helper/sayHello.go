package helper

import "fmt"

func SayHello(name string) { //jika di eksport tulis func dengan huruf besar di awal selain func seperti Variable juga bisa di buat dengan huruf besar lalu di ekport
	fmt.Println("Hello", name)
}