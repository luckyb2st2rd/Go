package main

import (
	"fmt"
	"strings"
)

func main() {
	//Работа со строками
	s := "Привет, Golang!"

	fmt.Println(len(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.Contains(s, "Golang"))
	fmt.Println(strings.Replace(s, "Golang", "Python", 1))

	name := "Alex"
	age := 19
	if age <= 20 {
		msg := fmt.Sprintf("Меня зовут %s, мне %d лет", name, age)
		fmt.Println(msg)
	} else {
		msg := fmt.Sprintf("Меня зовут %s, мне %d года", name, age)
		fmt.Println(msg)
	}

	json := `{
	"name":"Alex",
	"age":19
}`
	fmt.Println(json)
}
