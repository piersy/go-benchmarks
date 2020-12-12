package main

import "testing"

func BenchmarkSwapWithTempVar(b *testing.B) {
	x := 999999
	y := 777777
	for n := 0; n < b.N; n++ {
		xx := x
		yy := y
		temp := xx
		xx = yy
		yy = temp
		if xx != 777777 || yy != 999999 {
			panic("swap failed")
		}
	}
}

func BenchmarkSwapWithXor(b *testing.B) {
	x := 999999
	y := 777777
	for n := 0; n < b.N; n++ {
		xx := x
		yy := y
		xx ^= yy
		yy ^= xx
		xx ^= yy
		if xx != 777777 || yy != 999999 {
			panic("swap failed")
		}
	}
}
