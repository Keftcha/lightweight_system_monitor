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

func uptime_and_average() (string, string) {
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
	uptime = fmt.Sprintf("Uptime: %s", uptime)

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

	average := fmt.Sprintf(
		"%d CPU(s) | Load average : %s, %s, %s",
		runtime.NumCPU(),
		strconv.FormatFloat(avg_1min, 'f', 2, 64),
		strconv.FormatFloat(avg_5min, 'f', 2, 64),
		strconv.FormatFloat(avg_15min, 'f', 2, 64),
	)

	return uptime, average
}
