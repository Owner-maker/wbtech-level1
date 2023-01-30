package main

import (
	"errors"
	"fmt"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/

// Объявление константы возраста - порога взрослого человека
const adultAgeStep = 18

//Объявление новой структуры

type Human struct {
	name   string
	age    int
	height int
}

//Объявление новой структуры

type Action struct {
	// "встраиваем" поле типа Human -> аналог наследования, композиция
	// при этом поле Human является указателем, то есть при создании Action будем передавать
	// Human по ссылке, а не копией
	*Human
}

//Объявление функции для создания объекта типа Human

func newHuman(name string, age, height int) (*Human, error) {
	// используя условную конструкцию, проверяем валидность переданных в функцию будущих значений полей Human
	// если резальтат проверки true -> функция возвращает nil в качестве ссылки на Human и саму ошибку
	if name == "" {
		return nil, errors.New("name can not be empty")
	}
	if age <= 0 {
		return nil, errors.New("age can not be less or equals zero")
	}
	if height <= 0 {
		return nil, errors.New("height can not be less or equals zero")
	}

	// используя & возвращаем ссылку (адрес памяти) на объект типа Human, в качестве ошибки - nil
	return &Human{
		name:   name,
		age:    age,
		height: height,
	}, nil
}

// Объявление метода isAdult, возвращаяющая булево значение - является ли Human взрослым

func (h Human) isAdult() bool {
	return h.age > adultAgeStep
}

func main() {
	//Создаем переменную типа Human и инициализируем значениями, используя функцию NewHuman
	human, err := newHuman("Petr", 30, 160)
	if err != nil {
		return
	}

	//Создаем переменную типа Action и присваиваем в поле типа Human значение типа Human, созданного ранее
	action := Action{Human: human}

	//Вызываем метод isAdult поля Human, который является анонимным, и выводим результат в консоль
	fmt.Print(action.isAdult()) // вывод: true
}
