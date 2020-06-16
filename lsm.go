package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Curent date
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("\033[01;36m%s\033[0m\n", date)

	// Uptime
	fmt.Printf("Uptime: %s\n", formatedUptime())

	// Load average
	fmt.Printf(
		"%d CPU(s) | Load average: %s\n",
		runtime.NumCPU(),
		formatedLoads(runtime.NumCPU()),
	)
}
