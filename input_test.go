package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadHostList(t *testing.T) {
	in := strings.NewReader("example.com\n\n# a comment\ngithub.com:443\n  cloudflare.com  \n")
	got := readHostList(in)
	want := []string{"example.com", "github.com:443", "cloudflare.com"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
