package main

import (
	"fmt"
)

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

	/*v := "start"

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
	}*/
	/*for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i == rand.Intn(10) {
			break
		}
	}
	i := 0
	fmt.Println("\n")
	for {
		i++
		fmt.Println(i)
		if i == rand.Intn(35) {
			break
		}
	}*/

	/*N, sum, pol, otr := 7, 0, 0, 0
	for i := 0; i <= N; i++ {
		j := rand.Intn(100)
		fmt.Println("Ваше число ", i, ": ", j)
		if j%2 == 0 {
			pol = pol + j
		} else {
			otr = otr + j
		}
		sum = sum + j
	}

	fmt.Println("Сумма всех чисел:", sum)
	fmt.Println("Сумма четных чисел:", pol)
	fmt.Println("Сумма нечетных чисел:", otr)*/
	/*a, b := 2, 5
	fmt.Println("Сумма чисел равна:", hello(a, b))*/

	number := 10

	pointer := &number

	*pointer = 20

	fmt.Println(number)
	fmt.Println("n:", *pointer)
	fmt.Println(pointer)

	a, b := 5, 10

	swap(&a, &b)

	fmt.Println("a:", a, "\nb:", b)
}

func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	fmt.Println("a:", *a, "\nb:", *b)
}

/*func hello(a int, b int) int {
	defer fmt.Println("Функция завршена1")
	defer func() {
		fmt.Println("Функция завершена 2")
	}()
	fmt.Println("Привет мир!")
	return a + b
}*/
