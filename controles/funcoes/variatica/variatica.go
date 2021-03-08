package main

import "fmt"

func media(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0
	}
	total := .0
	for _, numero := range numeros {
		total += numero
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Println("Média", media(10, 15, 47, 9.8, 4.7, 87.7))
}
