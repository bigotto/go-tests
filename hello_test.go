package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bruno")
	want := "Hello, Bruno"

	if got != want {
		t.Errorf("got %q wuant %q", got, want)
	}

}