package main

import "fmt"

func main()  {

	score := 8
	switch score {

	case 0,1,2,3:
		fmt.Println("Terrible")
	case 4,5,6:
		fmt.Println("Mediocre")
	case 7,8,9:
		fmt.Println("Not bad")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score,"off the chart")





	}
}
