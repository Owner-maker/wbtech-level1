package main

import "sync"

/*
Реализовать конкурентную запись данных в map
--------------------------------------------
1 Способ: с использоанием мьютекса
*/

type Numbers struct {
	mu     sync.RWMutex
	values map[int]int
}

func (v *Numbers) Add(num int) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.values[num] = num
}

func main() {
	
}
