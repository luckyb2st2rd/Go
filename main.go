package main

import "fmt"

func main() {
	name := ""

	fmt.Println("Как тебя зовут?")
	fmt.Scanln(&name)

	fmt.Println("Меня зовут, " + name)
}
