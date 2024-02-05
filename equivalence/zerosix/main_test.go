package main

import (
	"testing"
)

func TestIsEquivalentToACopy(t *testing.T) {
	res, err := createResource("")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	res2, err := createResource("")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}

	if !res.Equivalent(res2).IsLocalHashEqual() {
		t.Fatalf("object is not localHashEqual to it's own copy")
	}

	if !res.Equivalent(res2).IsEquivalent() {
		t.Fatalf("object is not equivalent to it's own copy")
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

	if res.Equivalent(res2).IsLocalHashEqual() {
		t.Fatalf("object is localHashEqual to a different one")
	}
}
