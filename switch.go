package main

import "fmt"

func main() {
	name := "Zakariyah"

	switch name {
	case "Zakariyah":
		fmt.Println("Hello Zakariyah")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	name = "FirmanFirmanFirmannn"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
