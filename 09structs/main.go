package main

import "fmt"

func main() {
	sugam := User{"sugam", 20, true}

	fmt.Println(sugam)
	fmt.Printf("%+v\n", sugam)
	fmt.Println(sugam.Name, sugam.Age, sugam.Status)
}

type User struct {
	Name   string
	Age    int
	Status bool
}
