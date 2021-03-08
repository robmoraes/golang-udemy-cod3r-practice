package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[:2]

	fmt.Println(a, s)

	s[1] = 24
	fmt.Println(a, s)

	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)
	s1[3] = 67
	fmt.Println(s1, s2)
}
