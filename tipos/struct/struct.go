package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := produto{"Bola de basquete Adidas All Court Indoor/Outdoor", 157.98, 0.15}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{
		nome:     "Camisa Laker NBA",
		preco:    159.99,
		desconto: 0.10,
	}
	fmt.Println(produto2, produto2.precoComDesconto())
}
