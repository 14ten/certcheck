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

func issuerCN(c *x509.Certificate) string {
	if c.Issuer.CommonName != "" {
		return c.Issuer.CommonName
	}
	if len(c.Issuer.Organization) > 0 {
		return c.Issuer.Organization[0]
	}
	return fmt.Sprint(c.Issuer)
}

func check(host string, timeout time.Duration) Result {
	r := Result{Host: host}
	dialer := &net.Dialer{Timeout: timeout}
	conn, err := tls.DialWithDialer(dialer, "tcp", parseHost(host), &tls.Config{InsecureSkipVerify: true})
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

func checkAll(hosts []string, workers int, timeout time.Duration) []Result {
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
				results[i] = check(hosts[i], timeout)
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
