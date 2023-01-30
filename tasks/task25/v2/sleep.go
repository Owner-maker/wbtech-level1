package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.

2ой способ: с использованием условной конструкции и бесконечного цикла
*/

func sleep(d time.Duration) {
	startTime := time.Now()

	for {
		currTime := time.Now()                            // текущее время
		diff := currTime.Sub(startTime) / time.Nanosecond // разница между текущим и переданным функции в качестве параметра
		if diff >= d {                                    // если разница  равна или превышает переданного значения, значит отведенное время прошло, нужно остановить цикл для продолжения выполнения потока
			break
		}
	}
}

func main() {
	dur := time.Duration(5) * time.Second // задаем значение полю Duration в секундах нужного периода времени для ожидания
	fmt.Println(time.Now())
	sleep(dur)
	fmt.Println(time.Now())
}
