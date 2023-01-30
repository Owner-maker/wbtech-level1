package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	values := make(map[int][]float64) // мапа где ключ - диапазон (int), набор температур (слайс из float64)
	for i := 0; i < len(temps); i++ {

		// по ключу, полученному путем манипуляций с арифм. операциями, получаем круглое число - текущйи дипазаон
		// и добавляем значение текущее значение из слайса temps в внутренний слайс по ключу
		values[int(temps[i])/10*10] = append(values[int(temps[i])/10*10], temps[i])
	}

	// с помощью билдера, метода WriteString() создаем необходимую строку для удобного вывода на экран

	var builder strings.Builder
	for k, v := range values {
		builder.WriteString(strconv.Itoa(k))
		builder.WriteString(":{")
		for j := 0; j < len(v); j++ {
			builder.WriteString(fmt.Sprintf("%.1f", v[j]))
			if j+1 != len(v) {
				builder.WriteString(", ")
			}
		}
		builder.WriteString("}, ")
	}
	res := builder.String()
	res = res[0 : len(res)-2] // берем срез из строки, чтобы убрать последнюю запятую и пробел
	fmt.Print(res)
}
