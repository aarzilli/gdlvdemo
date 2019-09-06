package main

import "testing"

func TestAComplexFunction(t *testing.T) {
	for i := 0; i < 100; i++ {
		p, q := aComplexFunction(i)
		execute(p.X, q.Y)
	}
}
