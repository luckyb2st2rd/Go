package main

import "fmt"

// Переменные на уровне пакета (глобальные)
var globalVar string = "I`m global"

func main() {
	//Способ 1: полное объявление
	var name string = "Artem"

	//Способ 2: с выводом типа (type inference)
	var age = 25

	//Способ 3: короткое объявление := (только внутри функций!)
	city := "Stuttgart"

	//Объявление нескольких переменных
	var (
		x int     = 10
		y float64 = 3.14
		z bool    = true
	)

	//Множественное присваивание
	a, b := 1, 2
	a, b = b, a //Обмен значениями без временной переменной

	fmt.Println(name, city, age, x, y, z, a, b)

	//Константы и iota
	const Pi = 3.14159
	const Greeting = "Hello"
	fmt.Println(Pi, Greeting)

	//iota - автоинкремент для перечислений (enum)
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	//степени двойки через iota
	const (
		KB = 1 << (10 * (iota + 1)) //1024
		MB
		GB
	)
	fmt.Println(KB, MB, GB)
	fmt.Println(globalVar)
}
