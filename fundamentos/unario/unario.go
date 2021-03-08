package main

import "fmt"

func main() {
	x, y := 1, 2

	// apenas postfix
	x++ // x += 1 ou x = x+1

	y-- // y -= 1 ou y = y-1

	fmt.Println(x, y)

	// não pode:
	//    --x
	//    (x == y--)
}
