package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.

C использованием конструкции switch
*/

func printType(value interface{}) {
	// используя конструкцию switch и ключевое слово type динамически определяем тип данных входного параметра value
	switch v := value.(type) {
	// далее рассматриваем по отдельности каждый случай
	case int:
		fmt.Printf("type = \"%T\" and value = %v", v, v)
	case string:
		fmt.Printf("type = \"%T\" and value = %v", v, v)
	case bool:
		fmt.Printf("type = \"%T\" and value = %v", v, v)
	case chan int:
		fmt.Printf("type = \"%T\" and value = %v", v, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	printType(12)
}
