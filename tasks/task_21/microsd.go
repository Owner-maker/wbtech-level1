package main

import "fmt"

// объявление структуры MicroSd -> ее можно вставить, скопировать с нее информацию и вытащить, отключить
// все строки просто выводятся в консоль

type MicroSd struct {
}

func (m MicroSd) insert() {
	fmt.Println("Micro sd was successfully initialized")
}

func (m MicroSd) copyData() {
	fmt.Println("Data was copied to the PC")
}

func (m MicroSd) disconnect() {
	fmt.Println("Micro sd was successfully disconnected")
}
