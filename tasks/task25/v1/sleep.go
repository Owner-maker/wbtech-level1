package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.

1ый способ: с использованием функции NewTimer
*/

func sleep(d time.Duration) {
	t := time.NewTimer(d)
	if <-t.C; true { // выполняющий данный код поток блокируется, пока не получит значение текущего времени с потока
		return
	}
}

func main() {
	dur := time.Duration(5) * time.Second // задаем значение полю Duration в секундах нужного периода времени для ожидания
	fmt.Println(time.Now())
	sleep(dur)
	fmt.Println(time.Now())
}
