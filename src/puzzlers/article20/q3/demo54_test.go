package q3

import (
	"runtime"
	"testing"
)

func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}

func TestMaxProcs(t *testing.T) {
	println(runtime.GOMAXPROCS(0))
}
