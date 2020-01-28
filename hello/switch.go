package main

import (
	"fmt"
	"runtime"
)

func main() {
	for i := 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

	sistemOperasi := runtime.GOOS
	switch sistemOperasi {
	case "darwin":
		fmt.Println("Mac")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("Gak Tau")
	}
}
