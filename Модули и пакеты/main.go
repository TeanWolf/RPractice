package main

import (
	greeting "study/Greeting"
	"study/User"
)

func main() {
	greeting.SayHello()
	greeting.Bad()

	u := User.NewUser("Василий", 25)
	u.ChangeAge(23)
	u.GetInfo()
}
