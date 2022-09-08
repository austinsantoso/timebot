package main

import (
	"testing"
)

func TestOutputOne(t *testing.T) {
	output := 1
	value := outputOne()

	if output != value {
		t.Fail()
	}
	t.Fail()

}
