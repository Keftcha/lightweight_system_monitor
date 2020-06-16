package main

import (
	"testing"
)

func TestFormatAndUptimeWithLowUptimeAndLowAverage(t *testing.T) {
	uptimeAndAverage := " 12:21:19 up  3:29,  4 users,  load average: 0,00, 0,19, 0,46"
	expected_uptime := "Uptime: 3:29"
	expected_average := "4 CPU(s) | Load average: \033[35m0.00\033[0m, \033[35m0.19\033[0m, \033[35m0.46\033[0m"

	uptime_got, average_got := formatUptimeAndAverage(uptimeAndAverage, 4)

	if expected_uptime != uptime_got {
		t.Errorf("Expected: %s, but got: %s", expected_uptime, uptime_got)
	}

	if expected_average != average_got {
		t.Errorf("Expected: %s, but got: %s", expected_average, average_got)
	}
}

func TestFormatAndUptimeWithLowUptimeAndMultipleAverage(t *testing.T) {
	uptimeAndAverage := " 12:21:19 up  3:29,  23 users,  load average: 9,59, 0,19, 18,46"
	expected_uptime := "Uptime: 3:29"
	expected_average := "16 CPU(s) | Load average: \033[32m9.59\033[0m, \033[35m0.19\033[0m, \033[31m18.46\033[0m"

	uptime_got, average_got := formatUptimeAndAverage(uptimeAndAverage, 16)

	if expected_uptime != uptime_got {
		t.Errorf("Expected: %s, but got: %s", expected_uptime, uptime_got)
	}

	if expected_average != average_got {
		t.Errorf("Expected: %s, but got: %s", expected_average, average_got)
	}
}

func TestFormatAndUptimeWithOneDayUptimeAndVeryLowAverage(t *testing.T) {
	uptimeAndAverage := " 12:21:19 up  1 day, 17:52,  10 users,  load average: 0,02, 0,06, 0,08"
	expected_uptime := "Uptime: 1 day, 17:52"
	expected_average := "2 CPU(s) | Load average: \033[35m0.02\033[0m, \033[35m0.06\033[0m, \033[35m0.08\033[0m"

	uptime_got, average_got := formatUptimeAndAverage(uptimeAndAverage, 2)

	if expected_uptime != uptime_got {
		t.Errorf("Expected: %s, but got: %s", expected_uptime, uptime_got)
	}

	if expected_average != average_got {
		t.Errorf("Expected: %s, but got: %s", expected_average, average_got)
	}
}

func TestFormatAndUptimeWithTwoDaysUptimeAndLowAverage(t *testing.T) {
	uptimeAndAverage := " 12:52:34 up 2 days, 23:35,  2 users,  load average: 0,45, 0,22, 0,15"
	expected_uptime := "Uptime: 2 days, 23:35"
	expected_average := "8 CPU(s) | Load average: \033[35m0.45\033[0m, \033[35m0.22\033[0m, \033[35m0.15\033[0m"

	uptime_got, average_got := formatUptimeAndAverage(uptimeAndAverage, 8)

	if expected_uptime != uptime_got {
		t.Errorf("Expected: %s, but got: %s", expected_uptime, uptime_got)
	}

	if expected_average != average_got {
		t.Errorf("Expected: %s, but got: %s", expected_average, average_got)
	}
}

func TestFormatAndUptimeWithTwoDaysUptimeAndMediumAverage(t *testing.T) {
	uptimeAndAverage := " 10:55:33 up 2 days, 23:38,  0 users,  load average: 0.48, 0.32, 0.20"
	expected_uptime := "Uptime: 2 days, 23:38"
	expected_average := "1 CPU(s) | Load average: \033[36m0.48\033[0m, \033[34m0.32\033[0m, \033[34m0.20\033[0m"

	uptime_got, average_got := formatUptimeAndAverage(uptimeAndAverage, 1)

	if expected_uptime != uptime_got {
		t.Errorf("Expected: %s, but got: %s", expected_uptime, uptime_got)
	}

	if expected_average != average_got {
		t.Errorf("Expected: %s, but got: %s", expected_average, average_got)
	}
}

func TestFormatAndUptimeWithBigUptimeAndMediumAverage(t *testing.T) {
	uptimeAndAverage := " 08:11:22 up 146 days, 34 min,  3 users,  load average: 0.34, 0.65, 0.18"
	expected_uptime := "Uptime: 146 days, 34 min"
	expected_average := "1 CPU(s) | Load average: \033[36m0.34\033[0m, \033[32m0.65\033[0m, \033[34m0.18\033[0m"

	uptime_got, average_got := formatUptimeAndAverage(uptimeAndAverage, 1)

	if expected_uptime != uptime_got {
		t.Errorf("Expected: %s, but got: %s", expected_uptime, uptime_got)
	}

	if expected_average != average_got {
		t.Errorf("Expected: %s, but got: %s", expected_average, average_got)
	}
}
