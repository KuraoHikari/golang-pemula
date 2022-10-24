package main
import ("golang-pemula/database"
		"fmt")
//import ... "golang-pemula/database" // blank identifier

func main(){
	res := database.GetDatabase()
	fmt.Println(res)
}