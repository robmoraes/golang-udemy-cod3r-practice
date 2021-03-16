package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("Primeiro!")
	ch <- 2
	fmt.Println("Segundo!")
	ch <- 3
	fmt.Println("Terceiro!")
	ch <- 4
	fmt.Println("Quarto!")
	ch <- 5
	fmt.Println("Quinto!")
	ch <- 6
	fmt.Println("Sexto!")
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
