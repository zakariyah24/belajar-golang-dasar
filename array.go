package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Zakariyah"
	names[1] = "Firmansyah"
	names[2] = "Pantouw"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		95,
		100,
		110,
		110,
	}

	for _, v := range values {
		fmt.Println(v)
	}

	fmt.Println(len(values))
}
