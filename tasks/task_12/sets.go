package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func main() {
	strings := [...]string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{}) // мапа с ключом типа string и значение - пустая структура
	for _, v := range strings {
		set[v] = struct{}{}
	}

	for k, _ := range set {
		fmt.Println(k)
	}
}
