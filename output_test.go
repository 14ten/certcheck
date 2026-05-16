package main

import (
	"testing"
	"time"
)

func TestExitCode(t *testing.T) {
	now := time.Now()
	cases := []struct {
		name    string
		results []Result
		want    int
	}{
		{"all ok", []Result{{Host: "a", NotAfter: now.Add(90 * 24 * time.Hour), DaysLeft: 90}}, 0},
		{"warn only", []Result{{Host: "a", DaysLeft: 20}}, 1},
		{"crit beats warn", []Result{{Host: "a", DaysLeft: 20}, {Host: "b", DaysLeft: 3}}, 2},
		{"err beats crit", []Result{{Host: "a", DaysLeft: 3}, {Host: "b", Error: "boom"}}, 3},
	}
	for _, c := range cases {
		if got := exitCode(c.results, 30, 7); got != c.want {
			t.Errorf("%s: got %d want %d", c.name, got, c.want)
		}
	}
}
