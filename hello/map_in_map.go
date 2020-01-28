package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"1001": "Eko",
		"1002": "Joni",
		"1003": "Susanto",
	}

	fmt.Println(mahasiswa)

	mahasiswa2 := map[string]map[string]string{
		"1001": map[string]string{
			"name":    "Eko",
			"address": "Indonesia",
			"gender":  "Male",
		},
		"1002": map[string]string{
			"name":    "Joni",
			"address": "China",
			"gender":  "Male",
		},
		"1003": map[string]string{
			"name":    "Susanti",
			"address": "Jepang",
			"gender":  "Female",
		},
	}
	// fmt.Println(mahasiswa2["1001"])
	fmt.Println(mahasiswa2["1001"]["address"])
	// delete(mahasiswa2, "1002")
	// fmt.Println(mahasiswa2)
}
