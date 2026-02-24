package main

import "fmt"

type User struct {
	Name      string
	Age       int
	isPremium bool
}

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

	/*var a int = 23

	p := &a

	swap(p)
	println(a)*/

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

/*func swap(p *int) {
	println("a1:", *p)
	*p = 25
	println("a2:", *p)
}*/

/*func hello(a int, b int) int {
	defer fmt.Println("Функция завршена1")
	defer func() {
		fmt.Println("Функция завершена 2")
	}()
	fmt.Println("Привет мир!")
	return a + b
}*/
