package main

import (
	"fmt"
	"time"
)

func main() {
	// Curent date
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("\033[01;36m%s\033[0m\n", date)

	uptime, average := uptimeAndAverage()
	fmt.Println(uptime)
	fmt.Println(average)
}
