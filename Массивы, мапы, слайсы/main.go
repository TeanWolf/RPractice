package main

import (
	"study/exerciseArr"

	"github.com/k0kubun/pp"
)

type User struct {
	Name      string
	Age       int
	isPremium bool
}

func main() {
	//arr := [5]int{5, 23, 1, 2, 45}

	userArray := [3]User{
		User{
			Name:      "Александр",
			Age:       23,
			isPremium: true},
		User{
			Name:      "Василий",
			Age:       45,
			isPremium: false},
		User{
			Name:      "Олег",
			Age:       31,
			isPremium: true}}

	//pp.Println(userArray)

	for i := 0; i < len(userArray); i++ {
		pp.Println(userArray[i])
	}

	for i, v := range userArray {
		pp.Println(i+1, v)
	}

	for i, v := range userArray {
		if v.Age > 20 {
			userArray[i].Age += 1
		}
		pp.Println(i, v)
	}

	pp.Println(userArray)

	exerciseArr.First()

	exerciseArr.Second()
}
