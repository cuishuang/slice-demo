package main

import (
	"testing"
)

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noPrepare()
	}
}

func BenchmarkTest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prePare()
	}
}

func BenchmarkTest3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prePareAppend()
	}
}

func BenchmarkTest4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		prePareCopy()
	}
}
