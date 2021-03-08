package main

import "fmt"

type item struct {
	productID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {
	total := .0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			{1, 1, 100.00},
			{2, 2, 23.47},
			{11, 100, 1.99},
		},
	}

	fmt.Printf("O valor total do pedido é R$ %.2f", pedido.valorTotal())
}
