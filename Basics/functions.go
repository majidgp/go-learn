package main

import "fmt"

func add(x,y int)int  {
	
	return  x+y
	
}

func location(name,city string)(region,country string) {

	switch city {

	case "shahi", "sari":
		country = "mz-iran"
	case "la", "new york":
		country = "na-usa"
	default:
		country = "unknown"
	}

	return
}
func main()  {
	fmt.Println(add(47,31))
	region,country := location("majid","shahi")
	fmt.Printf("%s %s",region,country)
}
