package main

import (
	"fmt"
	"sync"
)

/*
Реализовать конкурентную запись данных в map
--------------------------------------------
1 Способ: с использоанием мьютекса
*/

//создаем структуру Numbers с мьютексом и мапой, где кл.ч и значений тип int

type Numbers struct {
	mu     sync.Mutex
	values map[int]int
}

// функция записи в мапу значения

func (v *Numbers) Add(num int) {
	v.mu.Lock() // блокировка мьютекса для обеспечения последовательности доступа горутины к мапе
	v.values[num] = num
	v.mu.Unlock() // освобождение мьютекса
}

func main() {
	nums := &Numbers{
		values: make(map[int]int, 1000),
	}

	// итеративный вызов горутины, записывающей конкурентно значение итератора в мапу

	for i := 0; i < 1000; i++ {
		go func(nums *Numbers, val int) {
			nums.Add(val)
		}(nums, i)
	}

	fmt.Print(nums.values)
}
