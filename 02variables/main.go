package main

import "fmt"

const LoginToken string = "sdcsdaca"
// captical first char L mean we are making it public
// := dose not 
func main()  {
  var username string =  "sugam"
  fmt.Println(username)
  fmt.Printf("Variable is  of type: %T \n", username)

  var website = "learncode.com"
  fmt.Println(website)
  fmt.Printf("Variable is of type: %T \n", website)
  
  var isLoggedIn bool =  false
  fmt.Println(isLoggedIn)
  fmt.Printf("Variable is  of type: %T \n", isLoggedIn)

  var smallval uint8 = 255 
  fmt.Println(smallval)
  fmt.Printf("Variable is  of type: %T \n", smallval)

  var integer int
  fmt.Println(integer) 
  fmt.Printf("type : %T \n", integer)
  
  numberOfUser := 300000
  fmt.Println(numberOfUser)

  fmt.Println(LoginToken)
}
