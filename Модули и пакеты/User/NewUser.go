package User

import "fmt"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) User {
	if name == "" {
		fmt.Println("Введено пустое имя!")
		return User{}
	}
	if age < 0 || age > 110 {
		fmt.Println("Введен неверный возраст!")
		return User{}
	}

	return User{
		name: name,
		age:  age,
	}
}
