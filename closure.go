package main

import "fmt"

func main() {
	counter := 0
	increement := func() {
		fmt.Println("increement")
		counter++
	}

	increement()
	increement()
	increement()
	fmt.Println(counter)
}
