package main

import "fmt"

type Us struct {
	firstname,lastname string
}

func (u *Us)Fullname()string  {
	return fmt.Sprintf("%s %s",u.firstname,u.lastname)
}
type Nus struct {
	Us
	playerid  int
}

func (p *Nus)Fullname()string  {
	return fmt.Sprintf("%s %s is %d",p.firstname,p.lastname,p.playerid)
}
type Nam interface {
	Fullname() string
}

func Get(n Nam) string {
    return fmt.Sprintf("%s",n.Fullname())
}

func main()  {
	x := &Us{firstname: "majid",lastname: "hp"}
	y := &Nus{}
	y.firstname = "babak"
	y.lastname = "kh"
	y.playerid = 157

	fmt.Println(Get(x))
	fmt.Println(Get(y))
}