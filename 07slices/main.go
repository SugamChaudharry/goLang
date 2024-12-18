package main

import (
	"fmt"
	"sort"
)

func main() {
  var arr = []string{"chaudhary", "chaudhary", "nill", "chaudharychaudhary"}
  fmt.Printf("type of arr %T\n", arr)
  arr = append(arr, "tebatiya", "khureja")
  fmt.Println(arr)

  arr = append(arr[1:])
  fmt.Println(arr)
  
  marks := make([]int, 4)

  marks[0] = 123
  marks[1] = 143
  marks[2] = 163
  marks[3] = 173
  // marks[4] = 777 // error

  marks = append(marks, 132,143,154)

  fmt.Println(marks)

  sort.Ints(marks)
  fmt.Println(marks)
  
  // how to remove value from slices

  var cources = []string{"reactjs", "python", "go", "js/ts", "rust"}
  fmt.Println(cources)
  cources = append(cources[:2], cources[2+1:]...)

  fmt.Println(cources)
}
