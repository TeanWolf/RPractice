package User

import (
	"github.com/k0kubun/pp"
)

func (user *User) ChangeName(n string) {
	if n != "" {
		user.name = n
	}
}

func (user *User) ChangeAge(a int) {
	if a > 0 && a < 110 {
		user.age = a
	}
}

func (user User) GetInfo() {
	pp.Println(user)
}
