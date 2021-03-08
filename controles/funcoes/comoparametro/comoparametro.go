package main

import "fmt"

func multiplicacao(a, b float64) float64 {
	return a * b
}

func exec(funcao func(float64, float64) float64, p1, p2 float64) float64 {
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplicacao, 7, 9)
	fmt.Println(resultado)
}
