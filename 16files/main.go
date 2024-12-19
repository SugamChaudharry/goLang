package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This is my content ~ sugam"

	file, err := os.Create("./myOsFile")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("lenght is: ", length)
	defer file.Close()

	readFile(file.Name())
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(databyte))
}
