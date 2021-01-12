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

func making()  {
	cities := make([]string,3)
	cities[0] = "shahi"
	cities[1] = "sari"
	cities[2] = "vresk"
	fmt.Println(cities)
	fmt.Println("--------------------")
}
func aplen()  {
	city := []string{"shahi"}
	city = append(city,"sari")
	fmt.Println(city)
	fmt.Println(len(city),cap(city))
	fmt.Println("---------------------")
}
func ranges()  {
	pow := []int{1,2,4,8,16,32,64,128}
	for i,v := range pow{
		fmt.Printf("2**%d = %d\n",i,v)

	}
	fmt.Println("--------------------------")
}

func bc()  {
	x := []int{1,2,3,4,5,6,7,8,9,10}
	for i := range x{
        if i == 8{
        	break
		}
		if i == 3{
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("-----------------------")

}

func mp()  {

	celebs := map[string]int{
		"majid" : 30,
		"babak" : 22,
		"nima"  : 29,
		"mamad" : 29,
	}
	fmt.Println(celebs)

}
func main()  {
	slice()
	myslices()
	making()
	aplen()
	ranges()
	bc()
	mp()

}
