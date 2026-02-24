package main

type User struct {
	Name      string
	Age       int
	isPremium bool
}

func main() {
	var a int = 23

	p := &a

	swap(p)
	println(a)
}

func swap(p *int) {
	println("a1:", *p)
	*p = 25
	println("a2:", *p)
}
