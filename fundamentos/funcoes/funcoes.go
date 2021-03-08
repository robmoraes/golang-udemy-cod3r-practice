package main

import "fmt"

// deve ser executada pelo terminal com $ go run *.go

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}

func main() {
	resultado := somar(3, 4)
	imprimir(resultado)
}
