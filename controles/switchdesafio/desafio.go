package main

import "fmt"

// Exercício, reescrever exemplo de controles/ifelseif com switch

func notaParaConceito(n float64) string {
	switch {
	case n >= 9:
		return "A"
	case n >= 8:
		return "B"
	case n >= 5:
		return "C"
	case n >= 3:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.6))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(4.2))
	fmt.Println(notaParaConceito(2.1))
}
