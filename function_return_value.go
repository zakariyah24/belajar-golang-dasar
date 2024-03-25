package main

import "fmt"

func getHello(firstName string, lastName string) string {
	hello := "Hello " + firstName + lastName
	return hello

}
func main() {
	result := getHello("Zakariyah", "Firmansyah")
	fmt.Println(result)
}
