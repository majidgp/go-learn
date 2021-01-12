package main

import "fmt"

func slice()  {
	p := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(p)
	fmt.Println("-------------------")
}
func myslices()  {
	m := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(m[1:5])
	fmt.Println(m[:8])
	fmt.Println(m[4:])
	fmt.Println(m[0:3])
	fmt.Println("-----------------------")
}





func main()  {
	slice()
	myslices()
}
