package main

import "testing"

func TestParseHost(t *testing.T) {
	cases := []struct {
		in   string
		port int
		want string
	}{
		{"example.com", 443, "example.com:443"},
		{"example.com:8443", 443, "example.com:8443"},
		{"1.2.3.4", 443, "1.2.3.4:443"},
		{"[::1]:443", 443, "[::1]:443"},
		{"example.com", 8443, "example.com:8443"},
		{"1.2.3.4", 636, "1.2.3.4:636"},
	}
	for _, c := range cases {
		if got := parseHost(c.in, c.port); got != c.want {
			t.Errorf("parseHost(%q, %d) = %q, want %q", c.in, c.port, got, c.want)
		}
	}
}
