package main

import ("fmt"
		"reflect")

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0 ; i < t.NumField() ; i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

type ContohLagi struct {
	Name string
	Description string
}
type ContohLagi2 struct {
	Name string `required:"true" max:"10"`
	Description string `required:"true" max:"10"`
}

func main(){
	sample := Sample{"Kurao"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"", "anjay"} // tidak tervalidasi karena tidak ada requirednya
	fmt.Println(IsValid(contoh))

	contoh2 := ContohLagi2{"aaa", ""} 
	fmt.Println(IsValid(contoh2))
}