package main

import "fmt"

type User struct {
	Name      string
	Age       int
	isPremium bool
}

func main() {

	user1 := NewUser(
		"Василий",
		23,
		true,
	)

	user2 := NewUser(
		"Александр",
		156,
		false,
	)

	user1.Struct()
	user2.Struct()

	user1.Premium()
	user2.Premium()

	user1.Struct()
	user2.Struct()

}

func NewUser(name string, //конструктор для проверки
	age int,
	isPremium bool,
) User {
	if age < 0 || age > 110 {
		fmt.Println("Возраст не прошел валидацию!")
		return User{}
	}

	if name == "" {
		fmt.Println("Имя не прошел валидацию!")
		return User{}
	}

	return User{
		Name:      name,
		Age:       age,
		isPremium: isPremium,
	}

}

func (user User) Struct() {
	if user.Name != "" {
		fmt.Println("Вывод пользователя")
		fmt.Println("Имя пользователя:", user.Name)
		fmt.Println("Возраст пользователя:", user.Age)
		fmt.Println("Премиум:", user.isPremium)
	}
}

func (user *User) Premium() { //метод и ресивер
	user.isPremium = true
}
