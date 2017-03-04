package battery

import "testing"

const (
	cap    = "/sys/class/power_supply/BAT1/capacity"
	status = "/sys/class/power_supply/BAT1/status"
)

func BenchmarkLevel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Level(cap, status)
	}
}
