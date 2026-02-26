package main

import (
	"fmt"
	"study/excercise"
)

type User struct {
	Name string
	Age  int
}

func main() {
	userArray := []User{
		User{
			Name: "Александр",
			Age:  24,
		},
		User{
			Name: "Василий",
			Age:  56,
		},
	}

	fmt.Println(len(userArray))
	fmt.Println(cap(userArray))

	userArray = append(userArray, User{
		Name: "Виталик",
		Age:  23,
	})

	fmt.Println(len(userArray))
	fmt.Println(cap(userArray))

	intSlice := make([]int, 0) //без выделенного размера
	//intSlice := make([]int, 0, 5)//выделен размер 5, чтобы процессор не тратил время на перевыделение массива

	intSlice = append(intSlice, 3, 4, 5, 1, 2)

	fmt.Println(intSlice)

	excercise.First()
	excercise.Second()
	excercise.Third()
}
