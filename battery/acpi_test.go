package battery

import "testing"

func BenchmarkAcpiLevel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AcpiLevel()
	}
}
