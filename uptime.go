package main

import (
	"fmt"
	"golang.org/x/sys/unix"
)

func uptime() int64 {
	var info unix.Sysinfo_t
	unix.Sysinfo(&info)

	return int64(info.Uptime)
}

func uptimeDays() int64 {
	return int64(uptime() / 60 / 60 / 24)
}

func uptimeHours() int64 {
	return int64((uptime() - uptimeDays()*24*60*60) / 60 / 60)
}

func uptimeMinutes() int64 {
	return int64((uptime() - (uptimeDays()*24*60*60 + uptimeHours()*60*60)) / 60)
}

func formatedUptime() string {
	days := uptimeDays()
	hours := uptimeHours()
	minutes := uptimeMinutes()

	// String day
	var sDay string
	if days == 1 {
		sDay = "1 day"
	} else {
		sDay = fmt.Sprintf("%d days", days)
	}

	// Formated uptime
	sUptime := fmt.Sprintf("%01d:%02d", hours, minutes)
	if days != 0 {
		sUptime = fmt.Sprintf("%s %s", sDay, sUptime)
	}

	return sUptime
}
