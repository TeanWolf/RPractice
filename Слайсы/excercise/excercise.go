package excercise

import (
	"fmt"
	"math/rand"
)

func First() {
	slice := make([]int, 5)

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func Second() {
	sliceA := []int{1, 2, 3, 4}

	sliceB := []int{}

	sliceB = append(sliceB, sliceA...)

	fmt.Println(len(sliceB), cap(sliceB))
	sliceA[0] = 4

	fmt.Println(sliceA)
	fmt.Println(sliceB)
}

func Third() {
	slice := []int{1, 2, 3, 4, 5}
	n := 10
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(100))
	}
	fmt.Println(slice)
}
