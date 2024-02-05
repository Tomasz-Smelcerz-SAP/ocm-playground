package main

import (
	"testing"
)

func TestIsEquivalent(t *testing.T) {
	res, err := createResource("")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	res2, err := createResource("")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	if !res.IsEquivalent(res2) {
		t.Fatalf("Result is not equivalent to it's own copy")
	}
}

func TestIsLocalHashEqualToADifferentObject(t *testing.T) {
	res, err := createResource("-foo")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	res2, err := createResource("-bar")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	if res.IsEquivalent(res2) {
		t.Fatalf("object is equal to a different one")
	}
}
