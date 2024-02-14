package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai")

	// array , slice, map
	names := []string{"Zaka", "Firmansyah", "Pantouw"}
	for _, v := range names {
		fmt.Println(v)
	}
	for index, v := range names {
		fmt.Println(index, " = ", v)
	}
}
