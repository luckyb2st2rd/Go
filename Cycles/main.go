package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	res := 0
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if a < b && a <= 100 && b <= 100 {
		for i := a; i < b; i++ {
			res += i
		}
		res += b
		fmt.Println(res)
	}
}
