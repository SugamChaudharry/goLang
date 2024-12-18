package main

import "fmt"

func main ()  {
  var name [4]string
  
  name[0] = "sugam"
  name[1] = "bobby"
  name[3] = "sugambobby"

  fmt.Println(name, len(name))
  
  var surname = [4]string{"chaudhary", "chaudhary", "nill", "chaudharychaudhary"}

  fmt.Println(surname, len(surname))
}
