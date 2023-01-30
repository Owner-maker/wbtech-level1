package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.

Для решения данной задачи был использован пакет
*/

// созданные ниже 4 функции принимают в качестве параметра два параметра типа *big.Int, что важно именно по указателю область памяти
// далее происходит объявление переменной типа big.Int как результирующей
// затем у этой переменной вызывается необходимый метод, который соответствует нужной нам арифметической операции и возвращается указатель на данное значение

func sum(a, b *big.Int) *big.Int {
	var res big.Int
	res.Add(a, b)
	return &res
}

func sub(a, b *big.Int) *big.Int {
	var res big.Int
	res.Sub(a, b)
	return &res
}

func multiply(a, b *big.Int) *big.Int {
	var res big.Int
	res.Mul(a, b)
	return &res
}

func division(a, b *big.Int) *big.Int {
	var res big.Int
	res.Quo(a, b)
	return &res
}

func main() {

	// с помощью функции new() создаем две переменных типа big.Int, а именно указатели на области памяти
	a := new(big.Int)
	b := new(big.Int)

	// с помощью метода SetString() производим запись нужного значения в десятично с.с.
	a.SetString("123123123123123123123123123123123123123123", 10)
	b.SetString("123123123123123123123123123123123123123", 10)

	// примеры вывода функций
	fmt.Printf("sum = %d\n", sum(a, b))
	fmt.Printf("subtraction = %d\n", sub(a, b))
	fmt.Printf("multiplication = %d\n", multiply(a, b))
	fmt.Printf("division = %d\n", division(a, b))
}
