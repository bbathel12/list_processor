package main

import (
	"fmt"
	"testing"
)

func Benchmark_forceMd5(b *testing.B) {

	for n := 0; n < b.N; n++ {

		forceMd5(fmt.Sprint(n))
	}
}
