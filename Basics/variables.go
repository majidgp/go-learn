package main

import "fmt"

//1
//var(
//	name  string
//	age   int
//	city  string
//)

//2
//var(
//	name,city string
//	age       int
//)

//3
//var name string
//var age  int
//var city string

//4
//var(
//	 name string = "majid"
//	 age  int    = 30
//	 city string = "shahi"
//)

//5
//var (
//	name = "majid"
//	age  = 30
//	city = "shahi"
//)

//6
//var(
//	name,age.city = "majid",30,"shahi"
//)

func main()  {
	name,city := "majid","shahi"
	age       := 30
	fmt.Printf("%s %d %s",name,age,city)
}