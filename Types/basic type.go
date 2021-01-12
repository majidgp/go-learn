package main

import "fmt"

var(
	fl   float64   = 24.5
	max  uint      = 1<<12 -7
	clmx complex64 = 2+3i
	bl   bool      = true
)

func main()  {
	f := "%T(%v)\n"
	fmt.Printf(f,fl,fl)
	fmt.Printf(f,max,max)
	fmt.Printf(f,clmx,clmx)
	fmt.Printf(f,bl,bl)
}