package main

import "testing"

func TestSuccess(t *testing.T) {
	// add some change
	if simple() != 1 {
		t.Fatal("should success")
	}
}

func TestFail(t *testing.T) {
	if simple() != 2 {
		t.Fatal("test failed")
	}
}
