package main

import "testing"

func BenchmarkCompareRows(b *testing.B) {
	t := &Terrain{terrain: make([]rune, 10000), lineLength: 100}
	for i:=0; i < b.N; i++ {
		compareRows(0, 1, t)
	}
}

func BenchmarkCompareCols(b *testing.B) {
	t := &Terrain{terrain: make([]rune, 10000), lineLength: 100}
	for i:=0; i < b.N; i++ {
		compareCols(0, 1, t)
	}
}
