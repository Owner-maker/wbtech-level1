package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.

----------------------------------------------------------------

Способ 2	: ипользуя т.н. опрашивание одной горутиной другую посредством использования канала
*/

func main() {
	ch := make(chan int)
	stop := make(chan struct{}) // дополнительный канал опроcник типа struct{} -> пустая структура с 0 полей, занимает 0 бит памяти

	go func() {
		for {
			select {
			case ch <- 1: // пишем просто значение 1 в канал
			case <-stop: // если в stop каналае появилось значение -> канал следует закрыть
				close(ch)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		stop <- struct{}{} // отправляем в канал пустую структуру -> сигнал о завершении горутины
	}()

	for i := range ch {
		fmt.Printf("got %d\n", i)
	}

}
