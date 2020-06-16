package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func uptimeAndAverage() (string, string) {
	// Load average and uptime
	cmd := exec.Command("uptime")
	var output bytes.Buffer // output of the command
	cmd.Stdout = &output
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return formatUptimeAndAverage(output.String(), runtime.NumCPU())
}

func formatUptimeAndAverage(output string, nbCPU int) (string, string) {
	// Given the ouput of the `uptime` linux command, we format it

	// Separate infos in the `uptime` command
	infos := strings.Split(output, ",  ")

	// Get the uptime
	uptime := strings.Trim(strings.Split(infos[0], " up ")[1], " ")
	uptime = fmt.Sprintf("Uptime: %s", uptime)

	// Get the load average
	average_values := strings.Split(strings.Split(infos[2], ": ")[1], ", ")
	// Format number as float for the average
	formated_average := make([]string, 3)
	for idx, elem := range average_values {
		// Convert to float64
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
		// Choose color and convert in string
		color := defineColor(
			0,
			float64(nbCPU),
			value,
		)

		formated_average[idx] = fmt.Sprintf(
			"%s%s\033[0m",
			color,
			strconv.FormatFloat(value, 'f', 2, 64),
		)
	}

	avg_1min, avg_5min, avg_15min := formated_average[0], formated_average[1], formated_average[2]

	average := fmt.Sprintf(
		"%d CPU(s) | Load average: %s, %s, %s",
		nbCPU,
		avg_1min,
		avg_5min,
		avg_15min,
	)

	return uptime, average
}
