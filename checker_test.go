package main

import "testing"

func TestParseHost(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"example.com", "example.com:443"},
		{"example.com:8443", "example.com:8443"},
		{"1.2.3.4", "1.2.3.4:443"},
		{"[::1]:443", "[::1]:443"},
	}
	for _, c := range cases {
		if got := parseHost(c.in); got != c.want {
			t.Errorf("parseHost(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
