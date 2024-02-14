package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Subang", "Jawa Timur", "Indonesia"}
	var address2 *Address = &address1 // Copy by reference, menggunakan &, address1 juga berganti

	address2.City = "Jombang"
	fmt.Println(address1) // tidak berubah
	fmt.Println(address2) // subang berubah jadi Jombang

	// * merubah value address1
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)

}
