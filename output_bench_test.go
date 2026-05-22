package main

import (
	"io"
	"testing"
	"time"
)

func BenchmarkWriteTable(b *testing.B) {
	now := time.Now()
	results := make([]Result, 100)
	for i := range results {
		results[i] = Result{
			Host:     "example.com",
			NotAfter: now.Add(time.Duration(i) * 24 * time.Hour),
			DaysLeft: i,
			Issuer:   "Let's Encrypt R3",
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeTable(io.Discard, results, 30, 7, false)
	}
}
