package main

import "fmt"

func main() {
  langs := make(map[string]string)
  langs["js"] = "javascript"
  langs["py"] = "python"
  langs["ru"] = "ruby"

  fmt.Println(langs, langs["js"])
  delete(langs, "ru")

  fmt.Println(langs) 
  

  // loops for map
  for key,value := range langs {
    fmt.Println(key, value)
  }
}
