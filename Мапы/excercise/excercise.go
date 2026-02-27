package excercise

import (
	"fmt"

	"github.com/k0kubun/pp"
)

func First() {

	car := map[int]string{
		1: "BMW",
		2: "Volvo",
		3: "Zeekr",
		4: "Audi",
	}

	fmt.Println(car)

	for k, v := range car {
		fmt.Println(k, v)
	}

	k, v := 23, "Nissan"

	add(car, k, v)
	pp.Println(car)

	as := 23
	p := &as
	test(p)
	fmt.Println(*p)
}

func add(car map[int]string, k int, v string) {
	car[k] = v
	car[5] = "Porshe"
}

func test(p *int) {
	*p = 21
}
