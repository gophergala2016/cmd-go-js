package main

import "testing"

func TestBasic(t *testing.T) {
	if 1+2 != 3 {
		t.Error("failed a basic test")
	}
}
