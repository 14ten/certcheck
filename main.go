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
		quiet    = flag.Bool("quiet", false, "suppress output, exit code only")
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

	hosts := flag.Args()
	if len(hosts) == 0 {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			hosts = readHostList(os.Stdin)
		}
	}
	if len(hosts) == 0 {
		usage()
		os.Exit(2)
	}

	results := checkAll(hosts, *workers, *timeout)
	switch {
	case *quiet:
		// no output
	case *jsonOut:
		_ = writeJSON(os.Stdout, results)
	default:
		writeTable(os.Stdout, results, *warnDays, *critDays)
	}
	os.Exit(exitCode(results, *warnDays, *critDays))
}

func usage() {
	fmt.Fprintf(os.Stderr, `certcheck %s — TLS certificate expiry checker

Usage:
  certcheck [flags] host[:port]...

Flags:
`, version)
	flag.PrintDefaults()
}
