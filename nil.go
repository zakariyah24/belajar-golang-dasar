package main

import "fmt"

// interface, function, map, slice, pointer, dan channe

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}

}
func main() {
	data := NewMap("zakariyah")
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
