package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "Zakariyah",
		"address": "Malang",
	}

	fmt.Println(person["sfasd"])

	book := make(map[string]string, 0)
	book["title"] = "Buku Golang"
	book["author"] = "Zaka"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
