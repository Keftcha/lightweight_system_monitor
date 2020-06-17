package main

import (
	"math"
)

func defineColor(min float64, max float64, value float64) string {
	// Choose value, purple is lower, red is higher

	// Define colors
	colors := []string{
		"\033[95m", // Purple
		"\033[94m", // Blue
		"\033[96m", // Cyan
		"\033[92m", // Green
		"\033[93m", // Yellow
		"\033[91m", // Red
	}

	// Move our interval on [0:`max` - `min`] instead of [`min`:`max`]
	delta := min * -1
	max += delta
	value += delta

	// Cross product to move our interval on [1:6] because we have 6 colors
	color_index := int(math.Min(math.Max((6*value/max), 0), 5))

	return colors[color_index]
}
