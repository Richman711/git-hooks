package main

import "testing"

func TestAdd(t *testing.T) {
	if got, want := World(1), "New World!"; got != want {
		t.Errorf("World method produced wrong result. expected: %s, got: %s", want, got)
	}
}
