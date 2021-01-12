package main

import (
	"fmt"
	"time"
)

type bootcamp struct {
	lat,lon float64
	date    time.Time
}

func main()  {
	event := bootcamp{
		lat : 34.215,
		lon : 57.671,
	}
	event.date = time.Now()
	fmt.Printf("Event on %s location(%f,%f)",event.date,event.lat,event.lon)

}