package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений

-----------------------------------------------------------------

Первое решение, небезопасное: при малом количестве элементов массива есть небольшая, но все же вероятность того, что
один поток успеет считать значение накопителя, после этого другой поток считает накопитель, увеличит и присвоит новое значение,
после чего изначальный поток присвоит уже старое значение, в результате чего итоговое значение накопителя будет ниже, чем нужно.
*/

func main() {
	// создаем массив типа int[5] и заполняем значениями
	values := [5]int{2, 4, 6, 8, 10}

	sum := 0 // накопитель

	// объявляем переменную типа WaitGroup, которая позволит дождаться выполнения всех горутин
	var wg sync.WaitGroup

	// задаем цикл для прохождения по values
	for _, val := range values {

		// увеличиваем счетчик на 1 перед запуском очередной горутины
		wg.Add(1)
		// запуск горутины (легковесный поток), принимает в качестве параметра - int
		go func(val int) {
			// с помощью оператора defer объявляем, что необходимо после завершения горутины уменьшить счетчик на 1
			defer wg.Done()
			sum += val * val
			// сразу после объявлеия анонимной функции вызываем горутину
		}(val)
	}
	//вызываем метод Wait(), который позволит дождаться выполнения всех горутин благодаря счетчику
	//когда счетчик = 0 -> программа завершится
	wg.Wait()
	fmt.Print(sum)
}
