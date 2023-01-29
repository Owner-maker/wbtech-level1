package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.

2ой способ: с использованием спецаильной синтаксической структуры Go
*/

func main() {
	a := 4
	b := 8

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
