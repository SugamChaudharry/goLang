package main

import "fmt"

func main() {
	a := 2

	if a%2 == 0 {
		fmt.Println("odd")
	} else if a%2 != 0 {
		fmt.Println("odd")
	} else {
		fmt.Println("god knows")
	}

	if a := 44; a < 50 {
		fmt.Println(a, " is less than 50")
	} else if a == 50 {
		fmt.Println(a, " is equal to 50 ")
	} else {
		fmt.Println(a, " is garter 50")
	}
}
