package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestWriteCSV(t *testing.T) {
	results := []Result{
		{
			Host:     "example.com",
			NotAfter: time.Date(2026, 8, 12, 0, 0, 0, 0, time.UTC),
			DaysLeft: 82,
			Issuer:   "DigiCert",
			Subject:  "example.com",
		},
		{
			Host:  "broken.test",
			Error: "timeout",
		},
	}
	var buf bytes.Buffer
	if err := writeCSV(&buf, results, 30, 7); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines (header + 2 rows), got %d", len(lines))
	}
	if lines[0] != "host,expires,days_left,status,issuer,error" {
		t.Errorf("unexpected header: %s", lines[0])
	}
	if !strings.HasPrefix(lines[1], "example.com,2026-08-12,82,OK,DigiCert,") {
		t.Errorf("unexpected row 1: %s", lines[1])
	}
	if !strings.Contains(lines[2], "broken.test") || !strings.Contains(lines[2], "ERR") {
		t.Errorf("unexpected row 2: %s", lines[2])
	}
}
