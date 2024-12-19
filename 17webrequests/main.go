package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://projects.100xdevs.com/tracks/auth-mern/Authentication-1"

func main() {
	fmt.Println("web request")
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T/n", res)

	defer res.Body.Close()

	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)

	fmt.Println(content)
}
