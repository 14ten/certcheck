package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"strings"
	"sync"
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

func friendlyErr(err error) string {
	msg := err.Error()
	switch {
	case strings.Contains(msg, "i/o timeout"):
		return "timeout"
	case strings.Contains(msg, "no such host"):
		return "dns: no such host"
	case strings.Contains(msg, "connection refused"):
		return "connection refused"
	case strings.Contains(msg, "tls: handshake failure"):
		return "tls handshake failed"
	}
	return msg
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

// Options controls how a check is performed.
type Options struct {
	Timeout  time.Duration
	SNI      string
	Insecure bool // skip certificate chain verification
}

func check(host string, timeout time.Duration) Result {
	return checkWith(host, Options{Timeout: timeout, Insecure: true})
}

// checkWith performs a TLS dial with the supplied options.
func checkWith(host string, opts Options) Result {
	r := Result{Host: host}
	addr := parseHost(host)
	cfg := &tls.Config{InsecureSkipVerify: opts.Insecure}
	if opts.SNI != "" {
		cfg.ServerName = opts.SNI
	}
	dialer := &net.Dialer{Timeout: opts.Timeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", addr, cfg)
	if err != nil {
		r.Error = friendlyErr(err)
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

func checkAll(hosts []string, workers int, opts Options) []Result {
	if workers < 1 {
		workers = 1
	}
	results := make([]Result, len(hosts))
	jobs := make(chan int, len(hosts))
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				results[i] = checkWith(hosts[i], opts)
			}
		}()
	}
	for i := range hosts {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return results
}
