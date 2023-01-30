package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.

3ий способ: с использованием функции After
*/

func sleep(d time.Duration) {
	if <-time.After(d); true { // эквивалент функции NewTimer, за исключением эффективности
		return
	}
}

func main() {
	dur := time.Duration(5) * time.Second // задаем значение полю Duration в секундах нужного периода времени для ожидания
	fmt.Println(time.Now())
	sleep(dur)
	fmt.Println(time.Now())
}
