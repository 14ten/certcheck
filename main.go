package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.2.0"

func main() {
	var (
		warnDays = flag.Int("warn-days", 30, "warn when cert expires within N days")
		critDays = flag.Int("crit-days", 7, "critical when cert expires within N days")
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

	results := make([]Result, 0, flag.NArg())
	for _, host := range flag.Args() {
		results = append(results, check(host))
	}
	writeTable(os.Stdout, results, *warnDays, *critDays)
}

func usage() {
	fmt.Fprintf(os.Stderr, `certcheck %s — TLS certificate expiry checker

Usage:
  certcheck [flags] host[:port]...

Flags:
`, version)
	flag.PrintDefaults()
}
