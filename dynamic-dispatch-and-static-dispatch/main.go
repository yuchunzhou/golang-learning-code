package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Email    string
	Password string
}

type Auth interface {
	Login()
}

func (u User) Login() {
	// fmt.Println("Login..")
}

func main() {
	user := User{
		Email:    "abc@def.com",
		Password: "111",
	}

	// static dispatch
	fmt.Println(reflect.TypeOf(user).Name()) // User
	start := time.Now()
	for range 10000000 {
		user.Login()
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

	var user_ Auth = User{
		Email:    "abc@def.com",
		Password: "111",
	}

	// dynamic dispatch
	// more slowly than static dispatch
	fmt.Println(reflect.TypeOf(user_).Name()) // User
	start = time.Now()
	for range 10000000 {
		user_.Login()
	}
	end = time.Now()
	fmt.Println(end.Sub(start))
}
