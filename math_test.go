package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(10, 10)
	if total != 20 {
		t.Errorf("Soma(10, 10) = %d; want 20", total)
	}
}