package exercise

import "fmt"

func First() {
	arr := [10]int{56, 2, 4, 2, 8, 3, 45, 0, 3, 1}
	sum(arr)
}

func sum(arr [10]int) {
	sum := 0
	for i, _ := range arr {
		sum += arr[i]
	}
	fmt.Println("Сумма массива:", sum)
}

func Second() {
	arr := [2]int{1, 2}
	a := len(arr)
	tmp := 0
	for i := 0; i < len(arr)/2; i++ {
		a--
		tmp = arr[i]
		arr[i] = arr[a]
		arr[a] = tmp
	}
	fmt.Println(arr)
}
