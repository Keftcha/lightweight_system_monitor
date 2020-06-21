package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"strconv"
	"strings"
)

func loads() [3]uint64 {
	var info unix.Sysinfo_t
	unix.Sysinfo(&info)

    loads := [3]uint64{
        uint64(info.Loads[0]),
        uint64(info.Loads[1]),
        uint64(info.Loads[2]),
    }

	return loads
}

func load1() float64 {
	return float64(loads()[0]) / float64(1<<16)
}

func load5() float64 {
	return float64(loads()[1]) / float64(1<<16)
}

func load15() float64 {
	return float64(loads()[2]) / float64(1<<16)
}

func formatedLoads(cpus int) string {
	// Loads values
	loads := []float64{load1(), load5(), load15()}
	// Loads string values
	sLoads := make([]string, 3)

	for idx, load := range loads {
		sLoads[idx] = fmt.Sprintf(
			"%s%s\033[0m",
			defineColor(0, float64(cpus), load),
			strconv.FormatFloat(load, 'f', 2, 64),
		)
	}

	return strings.Join(sLoads, ", ")
}
