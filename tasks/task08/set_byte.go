package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
-------------------------------------------------------------------------------------
необычный способ - через массив рун
*/

// константы - значения юникода для единицы и нуля

const (
	runeZero = 48
	runeOne  = 49
)

func main() {
	num := int64(4)                  // число, необходимое для изменения
	res, err := setBit(num, 3, true) // вызываем функцию, передаем исходное число, порядкоый номер бита для изменения и булево значения для вставки
	if err != nil {
		return
	}
	fmt.Println(res)
}

func setBit(num int64, order uint8, val bool) (int64, error) {
	if order > 64 {
		return num, errors.New("order must be <= 64")
	}

	res := fmt.Sprintf("%b", num) // конвертируем число в строку байтового представления
	runesInput := []rune(res)     // создаем слайс рун из строки

	runesOutput := make([]rune, 64) // создаем новый слайс с начальным размером в 64 (бита)

	copy(runesOutput[0:len(runesInput)], runesInput[:]) // копируем с исходного слайса значения в новый, где 64 элемента
	for i := 0; i < 64; i++ {
		if runesOutput[i] == 0 { // если встречается 0 -> создаем новый слайс -> обрезаем до этого момента, чтобы при конвертировании в int64 не было ошибок из-за номера юникода 0
			runesOutput = runesOutput[:i]
			break
		}
	}

	// вставка нужного значения
	if val {
		runesOutput[order-1] = runeOne
	} else {
		runesOutput[order-1] = runeZero
	}

	// превращаем слайс рун в строку
	s := string(runesOutput)

	// конвертируем строку в int64
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return num, err
	}
	return i, nil
}
