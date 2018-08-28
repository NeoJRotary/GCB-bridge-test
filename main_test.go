package main

import "testing"

func TestSuccess(t *testing.T) {
	if simple() != 1 {
		t.Fatal("should success")
	}
}

func TestFailed(t *testing.T) {
	if simple() != 1 {
		t.Fatal("test failed")
	}
}
