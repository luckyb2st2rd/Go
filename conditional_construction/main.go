package main

import (
	"fmt"
	"math"
)

func main() {
	//Условная конструкция if
	a := 6
	b := 8
	if a < b {
		fmt.Println("a меньше, чем b")
	}

	//If с краткой инструкцией
	x := 2
	n := 3
	lim := 10
	// В инструкции перед условием вычисляется значение v
	if v := math.Pow(float64(x), float64(n)); v < float64(lim) {
		fmt.Printf("Результат: %.2f меньше лимита\n", v)
	}

	//Условные конструкции else if и else
	if a < b {
		fmt.Println("а меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a ровно b")
	}

	//Оператор switch
	var i int
	fmt.Println("Введите число: ")
	fmt.Scanln(&i)
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
