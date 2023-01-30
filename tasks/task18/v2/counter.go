package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.

2ой способ: с использованием атомика
*/

type Counter struct {
	value int64
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	// в цикле запускаем горутину и позволяем каждой из 3 инкрементировать поле value счетчика по 5 раз
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				atomic.AddInt64(&counter.value, 1) // увеличиваем значение поля value струкуры Counter с использованием пакета atomic
			}
			wg.Done()
		}()
	}
	wg.Wait() // доиждаемся выполнения всех горутин
	fmt.Print(counter.value)
}
