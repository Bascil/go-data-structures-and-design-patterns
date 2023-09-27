package main

import "testing"

func TestSum(t *testing.T){ // inject testing.T pointer
	a := 5
	b := 4

	expected := 9
	res := sum(a, b)
	if res != expected {
		t.Errorf("Sum function doesnt sum correctly. %d + %d isnt %d", a, b, res)
	}
}

func TestMultiply(t *testing.T){
	
}