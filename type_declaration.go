package main

import "fmt"

func main() {

	type NoKtp string

	var kepZaka NoKtp = "zaka"

	var contoh string = "222222"
	var contohKtp NoKtp = NoKtp(contoh)

	fmt.Println(kepZaka)
	fmt.Println(contohKtp)
}
