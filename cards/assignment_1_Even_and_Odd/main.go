package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, j := range s {
		if j%2 == 0 {
			fmt.Println(j, " = Even")
		} else {
			fmt.Println(j, " = Odd")

		}

	}

}
