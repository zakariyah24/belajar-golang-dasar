package main

import "fmt"

func main() {
	name := "Firmansyah"

	if name == "Zakariyah" {
		fmt.Println("Hello Zakariyah")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
