package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.

1ый способ: с использованием мьютекса
*/

type Counter struct {
	value int
	mu    sync.Mutex
}

// метод, инкрементирующий значение поля value структуры Counter с блокировкой Мьютекса, т.е. только одна горутина
// иммет доступ к записи поля в моменте

func (c *Counter) increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup


	// в цикле запускаем горутину и позволяем каждой из 3 инкрементировать поле value счетчика по 5 раз
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 5; j++ {
				counter.increment()
			}
			wg.Done()
		}()
	}
	wg.Wait() // доиждаемся выполнения всех горутин
	fmt.Print(counter.value)
}
