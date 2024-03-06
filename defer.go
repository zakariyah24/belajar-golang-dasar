package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() //defer digunakan untuk running function setelah function ini selesai
	fmt.Println("run application")
}

func main() {
	runApplication()
}
