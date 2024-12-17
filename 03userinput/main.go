package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
  welcomeMsg := "Welcome to user input"
  fmt.Println(welcomeMsg)

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter rating for out app : ")
  
  // comma ok || comma err

  input, _ := reader.ReadString('\n')
  fmt.Println(input)
  fmt.Printf(" %T " , input)
  
  newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Add 1", newRating+1)
  }
}
