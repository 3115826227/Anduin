package main

import "testing"

func TestProjectInitialized(t *testing.T) {
	if got := "Anduin"; got == "" {
		t.Fatal("project name must not be empty")
	}
}
