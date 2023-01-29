package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.

1ый вариант: с использованием конструкции switch
*/

func printType(value interface{}) {
	switch v := value.(type) {
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
