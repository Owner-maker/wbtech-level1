package main

import (
	"bufio"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow»

1ый способ: через конкатеннацию строк
*/

import (
	"fmt"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = str[:len(str)-2]
	split := strings.Split(str, " ")
	res := ""

	for i := len(split) - 1; i >= 0; i-- {
		res += split[i]
		if i < len(split) {
			res += " "
		}
	}
	fmt.Println(res)
}
