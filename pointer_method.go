package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	zaka := Man{"Zaka"}
	zaka.Married()
	fmt.Println(zaka.Name)
}
