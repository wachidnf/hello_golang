package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		// if i%2 == 0 {
		// 	continue
		// }

		if i == 10 {
			break
		}

		fmt.Println(i)
	}
}
