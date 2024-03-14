package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// pass by value
	// address1 := Address{"Subang", "Jawa Timur", "Indonesia"}
	// address2 := address1 // Copy value

	// address2.City = "Jombang"
	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // subang berubah jadi Jombang

	// pass by pointer reference
	var address3 Address = Address{"Subang", "Jawa Timur", "Indonesia"}
	var address4 *Address = &address3 // Copy by reference, menggunakan &, address1 juga berganti

	address4.City = "Jombang"
	fmt.Println(address3) // tidak berubah
	fmt.Println(address4) // subang berubah jadi Jombang

}
