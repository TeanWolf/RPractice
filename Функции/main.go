package main

import (
	"fmt"
)

func main() {

	a, b := 2, 5
	fmt.Println("Сумма чисел равна:", hello(a, b))
}

func hello(a int, b int) int {
	defer fmt.Println("Функция завршена1")
	defer func() {
		fmt.Println("Функция завершена 2")
	}()
	fmt.Println("Привет мир!")
	return a + b
}
