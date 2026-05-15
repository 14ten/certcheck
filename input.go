package main

import (
	"bufio"
	"io"
	"strings"
)

func readHostList(r io.Reader) []string {
	var hosts []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		hosts = append(hosts, line)
	}
	return hosts
}
