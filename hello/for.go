package main

import "fmt"

func main() {
	//increment
	for i := 1; i <= 10; i++ {
		fmt.Println(i, "Hello World")
	}

	fmt.Println("Selesai")

	//decrement
	for i := 5; i >= 1; i-- {
		fmt.Println(i, "Hello GoLang")
	}

	fmt.Println("Selesai")
}
