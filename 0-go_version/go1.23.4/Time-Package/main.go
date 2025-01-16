package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a fixed time zone with UTC+5
	myZone := time.FixedZone("CustomZone", 5*3600) // 5*3600 seconds = UTC+5

	// Create a time in UTC
	nowUTC := time.Now().UTC()
	fmt.Println("Current time in UTC:", nowUTC)

	// Convert time to the custom fixed zone
	nowInMyZone := nowUTC.In(myZone)
	fmt.Println("Current time in CustomZone:", nowInMyZone)

	// Format time in the custom zone
	fmt.Println("Formatted time:", nowInMyZone.Format("2006-01-02 15:04:05 MST"))
}
