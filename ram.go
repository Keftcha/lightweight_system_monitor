package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/keftcha/meminfo"
)

var units []string = []string{"o", "Ko", "Mo", "Go", "To", "Po", "Eo", "Zo", "Yo"}

type ram struct {
	total     uint64
	used      uint64
	free      uint64
	available uint64
}

type swap struct {
	total uint64
	used  uint64
	free  uint64
}

func newRAM() ram {
	meinfo, err := meminfo.NewMeminfo()
	if err != nil {
		log.Fatal(err)
	}

	used := (meinfo.MemTotal - meinfo.MemFree - meinfo.Buffers - meinfo.Cached)

	return ram{
		total:     meinfo.MemTotal * 1024,
		used:      used * 1024,
		free:      meinfo.MemFree * 1024,
		available: meinfo.MemAvailable * 1024,
	}
}

func newSwap() swap {
	meinfo, err := meminfo.NewMeminfo()
	if err != nil {
		log.Fatal(err)
	}

	return swap{
		total: meinfo.SwapTotal * 1024,
		free:  meinfo.SwapFree * 1024,
	}
}

func chooseUnitIdx(value uint64) int {
	val := float64(value)

	idx := 0
	for val > 1024 && idx < len(units)-1 {
		idx++
		val /= 1024
	}

	return idx
}

func formatedRAM() string {
	ram := newRAM()

	legend := fmt.Sprintf(
		"%20s %11s %11s %11s",
		"Total",
		"Used",
		"Free",
		"Available",
	)

	values := fmt.Sprintf(
		"Mem: %15s %20s %20s %20s",
		formatUncolored(ram.total),
		formatColored(0, ram.total, ram.used),
		formatColored(ram.total, 0, ram.free),
		formatColored(ram.total, 0, ram.available),
	)

	return fmt.Sprintf("%s\n%s", legend, values)
}

func formatedSwap() string {
	swap := newSwap()

	legend := fmt.Sprintf(
		"%20s %11s %11s",
		"Total",
		"Used",
		"Free",
	)

	values := fmt.Sprintf(
		"Swap: %14s %20s %20s",
		formatUncolored(swap.total),
		formatColored(0, swap.total, swap.total-swap.free),
		formatColored(swap.total, 0, swap.free),
	)

	return fmt.Sprintf("%s\n%s", legend, values)
}

func humanRedable(value float64) float64 {
	return value / (math.Pow(1024, float64(chooseUnitIdx(uint64(value)))))
}

func formatUncolored(value uint64) string {
	return fmt.Sprintf(
		"%s%s",
		strconv.FormatFloat(humanRedable(float64(value)), 'f', 2, 64),
		units[chooseUnitIdx(value)],
	)
}

func formatColored(min, max, value uint64) string {
	return fmt.Sprintf(
		"%s%s%s\033[0m",
		defineColor(float64(min), float64(max), float64(value)),
		strconv.FormatFloat(humanRedable(float64(value)), 'f', 2, 64),
		units[chooseUnitIdx(value)],
	)
}
