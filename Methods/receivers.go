package main

import (
	"fmt"
	"math"
)

type Name struct {

	firstname,lastname string
}

type Vertex struct {
	x,y float64
}

func (s *Name)newN() string {
	return fmt.Sprintf("salam %s %s",s.firstname,s.lastname)
}

func (v *Vertex) Scale(f float64)  {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertex) Abs()float64  {
	return math.Sqrt(v.x*v.x+v.y*v.y)
}

func main()  {
	user := &Name{
		firstname: "majid",
		lastname: "hp",
	}
	fmt.Println(user.newN())

	v := &Vertex{3,8}
	v.Scale(7)
	fmt.Println(v,v.Abs())
}