package main

import "fmt"

func main() {
	slice1 := make([]string, 3)
	slice1[0] = "eko"
	slice1[1] = "joni"
	slice1[2] = "ahmed"

	slice2 := slice1[:]
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "Budi"
	fmt.Println(slice1)
	fmt.Println(slice2)

	fmt.Println()

	slice3 := append(slice1, "budi", "nugraha")
	fmt.Println(slice1)
	fmt.Println(slice3)

	slice4 := make([]string, 5)
	copy(slice4, slice1)

	fmt.Println(slice1)
	fmt.Println(slice4)
}
