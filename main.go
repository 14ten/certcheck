package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.1.0"

func main() {
	showVer := flag.Bool("version", false, "print version and exit")
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

	for _, host := range flag.Args() {
		r := check(host)
		if r.Error != "" {
			fmt.Printf("%s\tERR\t%s\n", r.Host, r.Error)
			continue
		}
		fmt.Printf("%s\t%s\n", r.Host, r.NotAfter.Format("2006-01-02"))
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
