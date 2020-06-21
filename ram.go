package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/Cloud-Foundations/Dominator/lib/meminfo"
	"golang.org/x/sys/unix"
)

var units []string = []string{"o", "Ko", "Mo", "Go", "To", "Po", "Eo", "Zo", "Yo"}

type ram struct {
	total     uint64
	free      uint64
	available uint64
}

type swap struct {
	total uint64
	free  uint64
}

func newRAM() ram {
	var info unix.Sysinfo_t
	unix.Sysinfo(&info)

	mem, err := meminfo.GetMemInfo()
	if err != nil {
		log.Fatal(err)
	}

	return ram{
		total:     uint64(info.Totalram),
		free:      uint64(info.Freeram),
		available: uint64(mem.Available),
	}
}

func newSwap() swap {
	var info unix.Sysinfo_t
	unix.Sysinfo(&info)

	return swap{
		total: uint64(info.Totalswap),
		free:  uint64(info.Freeswap),
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
		"%20s %11s %11s",
		"Total",
		"Free",
		"Available",
	)

	values := fmt.Sprintf(
		"Mem: %15s %20s %20s",
		formatUncolored(ram.total),
		formatColored(ram.total, 0, ram.free),
		formatColored(0, ram.total, ram.available),
	)

	return fmt.Sprintf("%s\n%s", legend, values)
}

func formatedSwap() string {
	swap := newSwap()

	legend := fmt.Sprintf(
		"%20s %11s",
		"Used",
		"Total",
	)

	values := fmt.Sprintf(
		"Swap: %23s %11s",
		formatColored(0, swap.total, swap.total-swap.free),
		formatUncolored(swap.total),
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
