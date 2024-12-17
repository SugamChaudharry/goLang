package main

import "fmt"

func main() {
  var ptr *int 
  
  fmt.Println(ptr)

  myNumber := 123
  
  var ptrr = &myNumber

  fmt.Println(myNumber, ptrr)
  
  *ptrr = *ptrr + 2

  fmt.Println(myNumber , ptrr)
}
