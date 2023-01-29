package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.

1ый способ: используя конкатенацию строк
*/

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(str)
	res := ""

	for i := len(runes) - 1; i >= 0; i-- {
		res += string(runes[i])
	}
	fmt.Print(res)
}
