package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.

3ий способ: с использованием XOR
*/

func main() {
	a := 4
	b := 8

	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a)
	fmt.Println(b)
}
