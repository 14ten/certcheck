package main

import (
	"errors"
	"testing"
)

func TestFriendlyErr(t *testing.T) {
	cases := []struct {
		name string
		in   error
		want string
	}{
		{"timeout", errors.New("dial tcp 1.2.3.4:443: i/o timeout"), "timeout"},
		{"dns", errors.New("dial tcp: lookup nope.invalid: no such host"), "dns: no such host"},
		{"refused", errors.New("dial tcp 127.0.0.1:1: connect: connection refused"), "connection refused"},
		{"handshake", errors.New("remote error: tls: handshake failure"), "tls handshake failed"},
		{"passthrough", errors.New("something exotic"), "something exotic"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := friendlyErr(c.in); got != c.want {
				t.Errorf("got %q want %q", got, c.want)
			}
		})
	}
}
