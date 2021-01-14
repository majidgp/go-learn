package main

import (
	"fmt"
	"strings"
)

type Mystr string
type Myfloat float64

func (s Mystr)Uppercase()string  {
	return strings.ToUpper(string(s))
}

func (f Myfloat) Abs() float64  {
	if f <0{
		return float64(-f)
	}
	return float64(f)
}

func main()  {

	s := Mystr("majid")
	f := Myfloat(-14)
	fmt.Println(s.Uppercase())
	fmt.Println(f.Abs())
}


