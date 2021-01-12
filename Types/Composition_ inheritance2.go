package main

import "fmt"

type User struct {
	name   string
	family string
	age    int
}

type Player struct {
	User
	id    uint64
}

func main()  {
	user := new(Player)
	user.name   = "majid"
	user.family = "hp"
	user.age    = 30
	user.id     = 241
	fmt.Println(user.id)
	fmt.Println(user.name)
	fmt.Println(user.family)
	fmt.Println(user.age)
}