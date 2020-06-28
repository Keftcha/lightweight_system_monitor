package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

var displayDate bool
var displayUptime bool
var displayAverage bool
var displayRAM bool
var displaySwap bool
var displayAll bool

func init() {
	flag.BoolVar(&displayDate, "date", false, "Display the durrent date with format YYYY-MM-DD HH-mm-SS")
	flag.BoolVar(&displayUptime, "uptime", false, "Display the uptime of the computer")
	flag.BoolVar(&displayAverage, "average", false, "Display the computer CPU load averag (1 min, 5 min, 15 min)")
	flag.BoolVar(&displayRAM, "ram", false, "Display the total ram, free ram and available ram")
	flag.BoolVar(&displaySwap, "swap", false, "Display the swap used and total swap space")
	flag.BoolVar(&displayAll, "all", false, "Display all information (date, uptime, CPU load average, RAM usage and swap usage")

	flag.Parse()

	// If no command line arguments are given, we display all
	if !displayDate && !displayUptime && !displayAverage && !displayRAM && !displaySwap {
		displayAll = true
	}
}

func main() {
	// Curent date
	if displayDate || displayAll {
		date := time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("\033[96m%s\033[0m\n", date)
	}

	// Uptime
	if displayUptime || displayAll {
		fmt.Printf("Uptime: \033[93m%s\033[0m\n", formatedUptime())
	}

	// Load average
	if displayAverage || displayAll {
		fmt.Printf(
			"%d CPU(s) | Load average: %s\n",
			runtime.NumCPU(),
			formatedLoads(runtime.NumCPU()),
		)
	}

	// Memory usage
	if displayRAM || displayAll {
		fmt.Println()
		fmt.Println(formatedRAM())
	}
	if displaySwap || displayAll {
		fmt.Println()
		fmt.Println(formatedSwap())
	}
}
