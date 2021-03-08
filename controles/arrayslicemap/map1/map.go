package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[78945678987] = "Ana"
	aprovados[98765432112] = "Pedrinho"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d\n", nome, cpf)
	}

	delete(aprovados, 12345678978)
	fmt.Println(aprovados)
}
