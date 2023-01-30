package main

import (
	"errors"
	"fmt"
)

/*
Удалить i-ый элемент из слайса.

1ый способ:
преимущество: сохраняется порядок элементов
недостаток: происходит копирование элементов массива для осуществления сдвига. Время O(n)
*/

func deleteElementInSlice(arr []int, i int) ([]int, error) {
	if i < 0 || i >= len(arr) {
		return nil, errors.New("index can not be < 0 or >= then length of array")
	}

	copy(arr[i:], arr[i+1:]) // копируем элементы слайса по правую сторону от i-ого элемента включительно в новый слайс, со всем элементами по правую сторону от i-ого элемента
	arr[len(arr)-1] = 0      // затираем последний элемент значением 0 <- по умолчанию для типа int
	arr = arr[:len(arr)-1]   // создаем новый слайс, "отрезая" последний затертый элемент
	return arr, nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res, err := deleteElementInSlice(arr, 2)
	if err != nil {
		return
	}
	fmt.Print(res) // вывод -> [1 2 5 4]
}
