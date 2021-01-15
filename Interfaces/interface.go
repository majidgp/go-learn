package main

import "fmt"

type User struct {
	firstname,lastname string
}

func (s *User)Name() string  {
	return fmt.Sprintf("%s %s",s.firstname,s.lastname)
}

type Namer interface {
	Name() string
}

func inter(n Namer)string  {
	return fmt.Sprintf("%s",n.Name())
}
func main()  {
	u := &User{firstname: "majid",lastname: "hp"}
	fmt.Println(inter(u))
}
