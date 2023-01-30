package main

import (
	"errors"
	"fmt"
)

/*
Удалить i-ый элемент из слайса.

1ый способ:
преимущество: скорость, нет необходимости перезаписывать и копировать все значения для сдвига, значение текущего элементу по индексу i
меняется на значение последнего элемента. Константное время
недостаток: не сохраняется порядок элементов в массиве
*/

func deleteElementInSlice(arr []int, i int) ([]int, error) {
	if i < 0 || i >= len(arr) {
		return nil, errors.New("index can not be < 0 or >= then length of array")
	}

	arr[i] = arr[len(arr)-1] // копируем значение последнего элемента в iый элемент
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
