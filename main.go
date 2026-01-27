package main

import "fmt"

func main() {
	/*var a1 int = 52
	var a2 string = "FALSE"
	var a3 float64 = 5.3
	var a4 bool = true

	a1 = 23
	a2 = "TRUE"
	a3 = 5.8
	a4 = false

	fmt.Println("a1:", a1, " a2:", a2, " a3:", a3, " a4:", a4)*/

	v := "start"

	switch v {
	case "start":
		a1 := -2

		if a1 > 0 && a1 != 0 {
			fmt.Println("Число положительное!")
		} else if a1 < 0 && a1 != 0 {
			fmt.Println("Число отрицательное!")
		} else {
			fmt.Println("Число равно 0")
		}
	case "stop":
		fmt.Println("Приложение остановлено!")
	case "exit":
		fmt.Println("Выход из приложения! ")
		break
	default:
		break
	}
}
