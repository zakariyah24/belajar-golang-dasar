package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is ", customer.Name)

}

func main() {
	var jak Customer
	jak.Name = "Zakariyah Firmansyah"
	jak.Address = "Indonesia"
	jak.Age = 24

	fmt.Println(jak)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Indonesia", 26}
	fmt.Println(budi)

	budi.sayHello("Agus")
}
