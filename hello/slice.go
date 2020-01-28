package main

import "fmt"

func main() {
	names := []string{
		"Eko",
		"Budi",
		"Joko",
		"Rudi",
		"Brian",
	}

	fmt.Println(names)
	slice1 := names[1:4]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)

	names[1] = "Nugraha"
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println()
	slice1[1] = "Moro"
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
