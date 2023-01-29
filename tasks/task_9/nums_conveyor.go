package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

const (
	n = 10
)

func main() {
	arr := [n]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inChan := make(chan int, n)  // канал дли "перезаписи" с массива arr
	outChan := make(chan int, n) // канал для записи
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		for i := 0; i < n; i++ {
			inChan <- arr[i]
		}
		wg.Done()
		close(inChan)
	}()

	go func() {
		for i := range inChan {
			outChan <- i * i
		}
		wg.Done()
		close(outChan)
	}()

	for i := range outChan {
		fmt.Println(i)
	}
	wg.Wait()
}
