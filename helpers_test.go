package piimasker_test

import (
	"testing"

	piimasker "piimasker"
)

// newTestMasker creates a Masker with default config for use in tests and benchmarks.
func newTestMasker(tb testing.TB) piimasker.PiiMasker {
	tb.Helper()
	return piimasker.NewMasker(piimasker.MaskerConfig{})
}

// newBenchMasker creates a Masker with default config for use in benchmarks.
func newBenchMasker(b *testing.B) piimasker.PiiMasker {
	b.Helper()
	return piimasker.NewMasker(piimasker.MaskerConfig{})
}

// runBench resets the benchmark timer, runs fn b.N times, and reports
// allocations per operation via b.ReportAllocs().
func runBench(b *testing.B, fn func()) {
	b.Helper()
	b.ReportAllocs()
	b.ResetTimer()
	for range b.N {
		fn()
	}
}
