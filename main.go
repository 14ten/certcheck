package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var (
		warnDays = flag.Int("warn-days", 30, "warn when cert expires within N days")
		critDays = flag.Int("crit-days", 7, "critical when cert expires within N days")
		jsonOut  = flag.Bool("json", false, "emit JSON instead of a table")
		quiet    = flag.Bool("quiet", false, "suppress output, exit code only")
		workers  = flag.Int("workers", 8, "concurrent checks")
		timeout  = flag.Duration("timeout", defaultTimeout, "per-host TLS dial timeout")
		port     = flag.Int("port", 443, "default TLS port when not specified per host")
		sni      = flag.String("sni", "", "override SNI server name (default: host)")
		insecure = flag.Bool("insecure", true, "skip cert chain verification (still reads expiry)")
		verbose  = flag.Bool("v", false, "verbose: log each host as it's checked")
		noColor  = flag.Bool("no-color", false, "disable ANSI colors")
		showVer  = flag.Bool("version", false, "print version and exit")
	)
	flag.Usage = usage
	flag.Parse()

	if *showVer {
		fmt.Println(fullVersion())
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

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lmicroseconds)
		log.SetOutput(os.Stderr)
		log.Printf("checking %d host(s) with %d worker(s), timeout=%s", len(hosts), *workers, *timeout)
	}
	results := checkAll(hosts, *workers, Options{
		Timeout:     *timeout,
		SNI:         *sni,
		Insecure:    *insecure,
		DefaultPort: *port,
	})
	if *verbose {
		log.Printf("done, %d result(s)", len(results))
	}
	sortByExpiry(results)
	switch {
	case *quiet:
		// no output
	case *jsonOut:
		_ = writeJSON(os.Stdout, results)
	default:
		writeTable(os.Stdout, results, *warnDays, *critDays, !*noColor && defaultColorEnabled())
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
