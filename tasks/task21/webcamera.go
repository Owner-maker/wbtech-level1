package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// объявление структуры WebCamera -> ее можно подключить, сделать фото и отключить
// все строки просто выводятся в консоль

type WebCamera struct {
}

func (w WebCamera) connect() {
	fmt.Println("Web camera was successfully connected")
}

func (w WebCamera) takePhoto() {
	fmt.Println("Photo was taken")
}

func (w WebCamera) disconnect() {
	fmt.Println("Web camera was successfully disconnected")
}
