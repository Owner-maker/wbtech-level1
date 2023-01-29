package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.

2ой способ: через string builder
*/

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(str)
	var builder strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteString(string(runes[i]))
	}
	fmt.Print(builder.String())
}
