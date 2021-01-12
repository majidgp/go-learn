package main

import "fmt"

const (
	pi           = 3.14
	statusok     = 200
	statuscreate = 201
	number       = 7
)

func main()  {

	name := "majid"
	fmt.Println(pi)
	fmt.Println(statusok)
	fmt.Println(statuscreate)
	c := 57/number
	fmt.Println(c)
	fmt.Printf("%s %d",name,number)
}