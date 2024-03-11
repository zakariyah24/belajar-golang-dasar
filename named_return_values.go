package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Zakariyah"
	middleName = "Firmansyah"
	lastName = "Pantouw"
	return firstName, middleName, lastName
}
func main() {

	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
