package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://projects.100xdevs.com:8000/tracks/auth-mern/Authentication-1?coursename=reactjs"

func main() {
	fmt.Println("welcome to hendle url in golang")
	fmt.Println(myurl)

	// parsing
	res, _ := url.Parse(myurl)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qprams := res.Query()
	fmt.Printf("%T\n", qprams)
	fmt.Println(qprams["coursename"])
	for _, val := range qprams {
		fmt.Println(val)
	}
	partOfurl := &url.URL{
		Scheme:  "https",
		Host:    "loc.dev",
		Path:    "tutcss",
		RawPath: "user=sugam",
	}

	anotherUrl := partOfurl.String()
	fmt.Println(anotherUrl)
}
