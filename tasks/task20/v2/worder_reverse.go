package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»

2ой способ: через string builder
*/

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n') // считываем с консоли строку до спец. символа - окончания строки
	str = str[:len(str)-2]                               // удаляем 2 последних символа (/r - возврат каретки и /n - перенос строки
	split := strings.Split(str, " ")                     // разбиваем строку на срез строк с разделителем в виде пробела
	var builder strings.Builder                          // объявление переменной типа Builder

	// циклом проходимся по срезу строк справа налево, добавляя значение в билдер, если элемент не последний -> добавляем пробел
	for i := len(split) - 1; i >= 0; i-- {
		builder.WriteString(split[i])
		if i < len(split) {
			builder.WriteString(" ")
		}
	}
	fmt.Println(builder.String())
}
