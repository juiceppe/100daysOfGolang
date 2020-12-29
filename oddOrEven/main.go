package main

import "fmt"

func main() {
	v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(v)

	for i := range v {
		if i%2 == 0 {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, "is Odd")
		}
	}
}
