package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func isUnique(s string) bool {
	s = strings.ToLower(s)            // создаем новую строку с символами нижнего регистра, т.е. превращаем все в единый вид
	values := make(map[rune]struct{}) // создаем мапу, где ключ - руна, значение - пустая структура
	runes := []rune(s)                // создаем срез рун на основе переданной строки

	// итеративно проверяем если уже есть такая руна в качестве ключа в мапе, то сразу возвращаем false -> символ встречатеся мин. 2 раз
	// значит строка не состоит из уникальных элементов
	for _, v := range runes {
		if _, ok := values[v]; ok {
			return false
		}
		values[v] = struct{}{}
	}
	return true // если цикл успешно был пройден -> возвращаем true
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n') // считываем символы строки с консоли до символа переноса строки
	fmt.Print(isUnique(str))
}
