package main

import "os"

const (
	ansiReset  = "\x1b[0m"
	ansiGreen  = "\x1b[32m"
	ansiYellow = "\x1b[33m"
	ansiRed    = "\x1b[31m"
	ansiGray   = "\x1b[90m"
)

func colorize(s, color string, enabled bool) string {
	if !enabled {
		return s
	}
	return color + s + ansiReset
}

func defaultColorEnabled() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	stat, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) != 0
}

func statusColor(status string) string {
	switch status {
	case "OK":
		return ansiGreen
	case "WARN":
		return ansiYellow
	case "CRIT", "ERR":
		return ansiRed
	}
	return ansiGray
}
