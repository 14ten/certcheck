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
		fmt.Println("TODO check:", host)
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
