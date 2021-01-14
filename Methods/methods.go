package main

import "fmt"

func Test() string {
	return "hello"
}

type User struct {
	name,lastname string
}

func (u User)newU()string  {
	return fmt.Sprintf("%s %s",u.name,u.lastname)
}
func main()  {
	user := User{
		name: "majid",
		lastname: "hp",
	}
	fmt.Println(user.newU())
	fmt.Println(Test())
}