package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Println("F2:", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprint("F4: ", p1, " ", p2)
}

func f5() (string, string) {
	return "F5", "abcdef"
}

func main() {
	f1()
	f2("abc", "def")
	fmt.Println(f3())
	fmt.Println(f4("abc", "def"))
	fmt.Println(f5())
}
