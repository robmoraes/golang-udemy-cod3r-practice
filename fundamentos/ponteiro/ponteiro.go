package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i
	*p++ // desreferenciando (pegando o valor)
	i++

	// Na versão do curso o Go não tinha aritmética de ponteiros

	fmt.Println(i, &i, *p, p)
}
