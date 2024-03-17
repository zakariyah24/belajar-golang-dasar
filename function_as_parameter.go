package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}

}

func main() {
	sayHelloWithFilter("Zaka", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
