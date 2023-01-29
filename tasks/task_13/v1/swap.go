package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.

1ый способ: путем сложения и вычитания
*/

func main() {
	a := 4
	b := 8

	a += b
	b = a - b
	a -= b

	fmt.Println(a)
	fmt.Println(b)
}
