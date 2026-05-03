package main

import (
	"crypto/tls"
	"strings"
	"time"
)

const defaultTimeout = 5 * time.Second

type Result struct {
	Host     string
	NotAfter time.Time
	Error    string
}

func parseHost(h string) string {
	if !strings.Contains(h, ":") {
		return h + ":443"
	}
	return h
}

func check(host string) Result {
	r := Result{Host: host}
	conn, err := tls.Dial("tcp", parseHost(host), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		r.Error = err.Error()
		return r
	}
	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		r.Error = "no peer certificates"
		return r
	}
	r.NotAfter = certs[0].NotAfter
	return r
}
