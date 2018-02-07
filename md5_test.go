package main

import (
	"testing"
)

func Benchmark_forceMd5(b *testing.B) {

	for n := 0; n < b.N; n++ {

		forceMd5(string(n))
	}
}
