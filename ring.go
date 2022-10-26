package main

import ("fmt"
		"container/ring"
		"strconv"
	)

func main(){
	var data *ring.Ring = ring.New(5)

	// data.Value = "Kurao"
	// var data2 = data.Next()
	// data2.Value = "Hikari"

	for i := 0; i < data.Len() ; i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	// fmt.Println(*data) //hasilnya {0xc0000543e0 0xc000054440 Data 0}

	data.Do(func(value interface{}){
		fmt.Println(value)
	})
}