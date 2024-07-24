package main

import (
	"fmt"
)

var a = 'G'

func mains() {
	a := [3]int{1, 2, 3}
	fmt.Printf("a: %T\n", a)

LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func n() {
	fmt.Println(a)
	fmt.Println(a)
	fmt.Println(a)
}

func m() {
	a := 'o'
	fmt.Println(a)
}
