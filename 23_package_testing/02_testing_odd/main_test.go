package main

import (
	"testing"
)

func TestOdd(t *testing.T) {
	if Odd(3) != 1 {
		t.Error("fail")

	}
}
