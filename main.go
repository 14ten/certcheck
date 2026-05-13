package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.3.0"

func main() {
	var (
		warnDays = flag.Int("warn-days", 30, "warn when cert expires within N days")
		critDays = flag.Int("crit-days", 7, "critical when cert expires within N days")
		jsonOut  = flag.Bool("json", false, "emit JSON instead of a table")
		workers  = flag.Int("workers", 8, "concurrent checks")
		timeout  = flag.Duration("timeout", defaultTimeout, "per-host TLS dial timeout")
		showVer  = flag.Bool("version", false, "print version and exit")
	)
	flag.Usage = usage
	flag.Parse()

	if *showVer {
		fmt.Println("certcheck", version)
		return
	}

	if flag.NArg() == 0 {
		usage()
		os.Exit(2)
	}

	results := checkAll(flag.Args(), *workers, *timeout)
	if *jsonOut {
		_ = writeJSON(os.Stdout, results)
	} else {
		writeTable(os.Stdout, results, *warnDays, *critDays)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `certcheck %s — TLS certificate expiry checker

Usage:
  certcheck [flags] host[:port]...

Flags:
`, version)
	flag.PrintDefaults()
}
