package main

import "fmt"

type User struct {
	Name     string
	Password string
}

type Auth interface {
	Login()
}

func (user User) Login() {
	fmt.Printf("user %s login with password: %s\n", user.Name, user.Password)
}

func main() {
	// static dispatch
	user := User{
		Name:     "aaa",
		Password: "111",
	}
	user.Login()

	// dynamic dispatch
	var user_ Auth = User{
		Name:     "bbb",
		Password: "222",
	}
	user_.Login()
}
