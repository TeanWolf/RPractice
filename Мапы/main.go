package main

import (
	"fmt"
	"study/excercise"

	"github.com/k0kubun/pp"
)

func main() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
		16: 0,
	}

	c, ok := weather[30]
	c1, ok1 := weather[16]

	fmt.Println(c, ok)
	fmt.Println(c1, ok1)

	weather[20] = +20
	fmt.Println(weather[20])

	for k, _ := range weather {
		weather[k] += 1
	}

	pp.Println(weather)

	criminal := map[string]bool{
		"Вася": true,
		"Петя": true,
		"Саша": false,
	}

	s := "Саша"
	c2, ok2 := criminal[s]

	if !ok2 {
		fmt.Println("Человека", s, "нет в базе")
	} else {
		fmt.Println("Человек", s, "найден в базе")
	}
	if !c2 {
		fmt.Println("Человек", s, "не судим")
	} else {
		fmt.Println("Человек", s, "судим")
	}

	excercise.First()

}
