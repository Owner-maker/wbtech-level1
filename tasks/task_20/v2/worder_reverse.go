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
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = str[:len(str)-2]
	split := strings.Split(str, " ")
	var builder strings.Builder

	for i := len(split) - 1; i >= 0; i-- {
		builder.WriteString(split[i])
		if i < len(split) {
			builder.WriteString(" ")
		}
	}
	fmt.Println(builder.String())
}
