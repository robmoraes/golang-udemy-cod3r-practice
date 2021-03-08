package main

import "fmt"

func closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}

	return funcao
}

func main() {
	x := 20

	imprimeX := closure()
	imprimeX() // x = 10 escopo da função closure

	fmt.Println(x) // x = 20 escopo da função main
}
