package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Matan")
	want := "Hello, Matan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
