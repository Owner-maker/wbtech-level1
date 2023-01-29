package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(a []int, el int) (int, bool) {
	left := 0            // индекс левой границы
	right := len(a) - 1  // индекс правой границы
	middle := len(a) / 2 // индекс середины

	// если искомый элемент уже превышает значение последнего эелемента или уже меньше самого первого эелемента, то выходим из функции
	if el > a[right] || el < a[left] {
		return 0, false
	}

	// цикл подолжается до тех пор, пока искомый элемент не будет равен элементу среза по индексу middle или пок левая граница <= правой
	for el != a[middle] && left <= right {

		// в нужном случае сдвигаем границу левую или правую
		if el < a[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
		middle = (right + left) / 2 // заново расчитываем индекс элемента посередине
	}

	// в зависимости от того нашли мы элемент или нет, выдаем ответ
	if a[middle] == el {
		return middle, true
	} else {
		return 0, false
	}
}

func main() {
	a := []int{2, 5, 7, 8, 10, 13, 16, 17}
	res, ok := binarySearch(a, 7)
	if ok {
		fmt.Print(res)
	} else {
		fmt.Println("Not found")
	}
}
