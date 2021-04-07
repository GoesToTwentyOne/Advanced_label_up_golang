package main

import "testing"

func TestCalc(t *testing.T) {
	if Calc(4) != 6 {
		t.Error("error")
	}

}
