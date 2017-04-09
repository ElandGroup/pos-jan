package main

import (
	"testing"
)

func TestDivision(t *testing.T) {
	if i, err := Division(6, 2); i != 3 || err != nil {
		t.Error("test zero is failure")
	} else {
		t.Log("test zero division is ok")
	}
}

func TestDivision2(t *testing.T) {
	t.Error("test is failure")
}

//1.go test
//2.go test -v
