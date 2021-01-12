package main

import "fmt"

func main()  {

	var a[3] string
	a[0] = "majid"
	a[1] = "hp"
	a[2] = "gp"
	fmt.Println(a)
	a[2] = "iran"
	fmt.Println(a)

	b := [2][3]string{{"this","is","first"},{"this","is","second"}}
	fmt.Println(b)
	b1 := [2][3]string{{"this","is","first"},{"this","is"}}
	fmt.Println(b1)

	c := [3][2]string{{"1","one"},{"2","two"},{"3","three"}}
	fmt.Println(c)
}
