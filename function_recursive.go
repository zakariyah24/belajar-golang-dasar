package main

import "fmt"

func factorialLoop(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}
func main() {
	fmt.Println(factorialLoop(10))
}
