package main

import "fmt"

func one()  {
	a := 0
	for i :=0;i<10;i++{
		a =+i
	}
	fmt.Println(a)
	fmt.Println("-------------")
}
func two()  {
	b := 1
	for ;b<10; {
		b =+1
	}
	fmt.Println(b)
	fmt.Println("----------")
}
func three()  {
	c := 1
	for c<10  {
		c =+1
	}
	fmt.Println(c)
	fmt.Println("----------")
}
func four()  {
	d := 1
	for  {
		d =+1
	}
	fmt.Println(d)
}
func main()  {

	one()
	two()
	three()
	four()
}
