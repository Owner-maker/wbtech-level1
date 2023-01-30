package main

import (
	"encoding/binary"
	"errors"
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
-------------------------------------------------------------------------------------
1 способ -
*/

func setByte(num int64, order uint8, val bool) (int64, error) {
	if order > 64 {
		return num, errors.New("order must be <= 64")
	}

	bytes := make([]byte, 8)
	// с помощью
	binary.LittleEndian.PutUint64(bytes, uint64(num)) // конвертирование в uint64 не приведет к сайд эффектам: бит знака сохранится

	bytes[order-1] = 0
	fmt.Println(bytes[order-1])
	fmt.Println(bytes)

	i := int64(binary.LittleEndian.Uint64(bytes))
	return i, nil

}

func main() {
	var num int64
	num = 3
	i, err := setByte(num, 2, false)
	if err != nil {
		return
	}
	fmt.Print(i)
}
