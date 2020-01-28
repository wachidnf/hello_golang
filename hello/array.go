package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Budi"
	names[1] = "Edi"
	names[2] = "Santoso"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, value := range names {
		// fmt.Println("index", index, "=", value)
		fmt.Println(value)
	}
}
