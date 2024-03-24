package main

import "fmt"

func getFullName(firstName string, lastName string) (string, string) {
	// hello := "Hello " + firstName + lastName
	return firstName, lastName

}
func main() {
	// first, last := getFullName("Zakariyah", "Firmansyah")
	// ignore first bisa menggunakan _
	_, last := getFullName("Zakariyah", "Firmansyah")
	fmt.Println(last)
}
