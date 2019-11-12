package txdata

import "testing"

func BenchmarkInitFunctionMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitFunctionMap()
	}
}
