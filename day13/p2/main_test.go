package main

import (
	"testing"
	"math/rand"
)

func BenchmarkCompareRows(b *testing.B) {
	t := &Terrain{terrain: make([]rune, 10000), lineLength: 100}
	for i:=0; i < b.N; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		compareRows(a, b, t)
	}
}

func BenchmarkCompareCols(b *testing.B) {
	t := &Terrain{terrain: make([]rune, 10000), lineLength: 100}
	for i:=0; i < b.N; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		compareCols(a, b, t)
	}
}
