package main

import "fmt"

func main() {
	Person := Person{
		Name:    "Vasya",
		Age:     25,
		Married: false,
		Hobbies: []string{
			"programming",
			"sport",
		},
		Address: Address{
			Country: "RU",
			City:    "Moscow",
		},
	}

	fmt.Println(Person)
}
