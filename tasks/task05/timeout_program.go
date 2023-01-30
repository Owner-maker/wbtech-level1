package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

// константы пределов для рандома

const (
	min = 0
	max = 100
)

func main() {
	// вызов метода Seed() из пакета rand -> позволяет каждый раз при вызове метода Intn() из пакета rand каждый раз новое псевдослучайное число
	rand.Seed(time.Now().UnixNano())

	// переменная типа int для хранения получаемой от пользователя величины
	var timeToWork int64

	_, err := fmt.Scan(&timeToWork) // считывание с консоли кол-во секунд работы программы
	if err != nil {
		return
	}

	// задание отложенной записи в переменную timer по истиченею времени timeToWork секунд
	timer := time.After(time.Duration(timeToWork) * time.Second)

	// создание буферизированного канала для основных значений, где размер буфера равен единице
	values := make(chan int, 1)
	defer close(values) // заранее объявляем о закрытии основного канала

	stop := make(chan bool, 1) // буф. канал с типом bool для сообщения горутине об остановке

	// задание и вызов горутины
	go func(values chan int) {
		for {
			select {
			// завершение горутины при получении сигнала об остановке, способ остановки горутины -> периодическое опрашивание
			case <-stop:
				fmt.Printf("reader goroutine has been stopped")
				return
				// дефолтная работа горутины - считывание значений с канала values
			default:
				fmt.Printf("value=%d\n", <-values)
			}
		}
	}(values)

	// основной цикл потока main
	for {
		select { // либо по умолчанию записывается новое значение в основной поток, либо при получении сигнала осн. поток останавливается
		case <-timer:
			fmt.Printf("main thread is stopped\n%d values are left in main channel", len(values))
			stop <- true
			return
			// дефолтная работа основного потока main - запись в канал values рандомных значений
		default:
			values <- rand.Intn(max-min+1) + min
			time.Sleep(time.Second)
		}
	}

}
