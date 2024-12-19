package main

import "fmt"

func main() {
	defer fmt.Println("5")
	defer fmt.Println("4")
	defer fmt.Println("3")
	defer fmt.Println("2")
	fmt.Println("1")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
