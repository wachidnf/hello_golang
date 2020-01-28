package main

import "fmt"

func main() {
	mahasiswa := make(map[string]string)

	mahasiswa["1001"] = "Eko"
	mahasiswa["1002"] = "Joni"
	mahasiswa["1003"] = "Susanto"

	fmt.Println(mahasiswa)

	for nim, name := range mahasiswa{
		fmt.Println("NIM", nim, "Nama:", name)
	}
}
