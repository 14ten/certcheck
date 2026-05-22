package main

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"testing"
)

func TestIssuerCN(t *testing.T) {
	cases := []struct {
		name string
		in   pkix.Name
		want string
	}{
		{
			"common name set",
			pkix.Name{CommonName: "Let's Encrypt R3", Organization: []string{"Let's Encrypt"}},
			"Let's Encrypt R3",
		},
		{
			"fallback to first org",
			pkix.Name{Organization: []string{"DigiCert Inc", "Other"}},
			"DigiCert Inc",
		},
		{
			"no cn or org",
			pkix.Name{Country: []string{"US"}},
			"C=[US]",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cert := &x509.Certificate{Issuer: c.in}
			got := issuerCN(cert)
			if c.name == "no cn or org" {
				if got == "" {
					t.Errorf("expected non-empty fallback, got empty")
				}
				return
			}
			if got != c.want {
				t.Errorf("got %q want %q", got, c.want)
			}
		})
	}
}
