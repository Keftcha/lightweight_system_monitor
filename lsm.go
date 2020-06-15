package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Curent date
	date := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("\033[01;36m%s\033[0m\n", date)

	// Load average and uptime
	cmd := exec.Command("uptime")
	var output bytes.Buffer // output of the command
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	infos := strings.Split(output.String(), ",  ")
	uptime := strings.Split(infos[0], " up ")[1]
	fmt.Printf("Uptime: %s\n", uptime)

	average_values := strings.Split(strings.Split(infos[2], ": ")[1], ", ")
	formated_average := make([]float64, 3)
	for idx, elem := range average_values {
		value, err := strconv.ParseFloat(
			strings.TrimSuffix(
				strings.ReplaceAll(elem, ",", "."),
				"\n",
			),
			64,
		)
		if err != nil {
			log.Fatal(err)
		}
		formated_average[idx] = value
	}

	avg_1min, avg_5min, avg_15min := formated_average[0], formated_average[1], formated_average[2]

	fmt.Printf(
		"%d CPU(s) | Load average : %s, %s, %s\n",
		runtime.NumCPU(),
		strconv.FormatFloat(avg_1min, 'f', 2, 64),
		strconv.FormatFloat(avg_5min, 'f', 2, 64),
		strconv.FormatFloat(avg_15min, 'f', 2, 64),
	)

}
