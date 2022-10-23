package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// func ChangeCountryToIndonesia(address Address){
// 	address.Country = "Indonesia"
// }
func ChangeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main(){
	address1 := Address {"Subang", "Jawa Barat", "Indonesia"}
	//address2 := address1 // reference btyvalue
	// address2 := &address1 // pass by reference
	// address2.City = "Bali"

	var address2 *Address = &address1

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1) // address 1 tidak berubah
	// fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // address 1  berubah
	fmt.Println(address2)

	// fmt.Println(address1)
	// fmt.Println(address2)

	// var address4 *Address = new(Address)
	var address4 *Address = &Address{"Subang","Kurao", "hikari"}
	address4.City = "Anjay"
	fmt.Println(address4)

	// var alamat = Address{
	// 	City: "Subang",
	// 	Province: "JawaBarat",
	// 	Country: "",
	// }
	// ChangeCountryToIndonesia(alamat)
	// fmt.Println(alamat) // tidak merubah value Address country tidak punya indonesia dan uncoment lanen 9-11

	var alamat = Address{
		City: "Subang",
		Province: "JawaBarat",
		Country: "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)// merubah value Address country jadi punya indonesia

	var alamatPointer = Address{
		City: "Bali",
		Province: "JawaBarat",
		Country: "",
	}
	var alamatPointerBali *Address = &alamatPointer
	ChangeCountryToIndonesia(alamatPointerBali)
	fmt.Println(alamatPointer)// merubah value Address country jadi punya indonesia
}