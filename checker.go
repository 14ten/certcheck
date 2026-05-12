package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"strings"
	"time"
)

const defaultTimeout = 5 * time.Second

type Result struct {
	Host     string    `json:"host"`
	NotAfter time.Time `json:"not_after,omitempty"`
	DaysLeft int       `json:"days_left,omitempty"`
	Issuer   string    `json:"issuer,omitempty"`
	Subject  string    `json:"subject,omitempty"`
	Error    string    `json:"error,omitempty"`
}

func parseHost(h string) string {
	if !strings.Contains(h, ":") {
		return h + ":443"
	}
	return h
}

func issuerCN(c *x509.Certificate) string {
	if c.Issuer.CommonName != "" {
		return c.Issuer.CommonName
	}
	if len(c.Issuer.Organization) > 0 {
		return c.Issuer.Organization[0]
	}
	return fmt.Sprint(c.Issuer)
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
	leaf := certs[0]
	r.NotAfter = leaf.NotAfter
	r.DaysLeft = int(time.Until(leaf.NotAfter).Hours() / 24)
	r.Issuer = issuerCN(leaf)
	r.Subject = leaf.Subject.CommonName
	return r
}
