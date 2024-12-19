package main

import "fmt"

func main() {
	days := []string{"mon", "tue", "web", "thu", "fri", "sat", "sun"}

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	fmt.Println("========================================")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("========================================")

	for i, day := range days {
		fmt.Println(i, day)
	}

	fmt.Println("========================================")

	j := 1

	for j < 10 {

		if j == 5 {
			break
		}

		if j == 4 {
			goto hii
		}

		fmt.Println(j)
		j++
	}

hii:
	fmt.Println("hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")

}
