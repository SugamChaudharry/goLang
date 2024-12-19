package main

import "fmt"

func main() {
	sugam := User{"sugam", 20, true}

	fmt.Println(sugam)
	fmt.Printf("%+v\n", sugam)
	fmt.Println(sugam.Name, sugam.Age, sugam.Status)

	sugam.GetStatus()
	sugam.NewStatus()
	sugam.NewStatus()
	sugam.NewStatus()
	sugam.NewStatus()
	sugam.NewStatus()
	sugam.NewStatus()
	sugam.NewStatus()
	sugam.NewStatus()

	fmt.Println(sugam.Name, sugam.Age, sugam.Status)
}

type User struct {
	Name   string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Status is: ", u.Status)
}

func (u User) NewStatus() {
	// u is copy of user not user struct
	u.Status = !u.Status
	fmt.Println("new status: ", u.Status)
}
