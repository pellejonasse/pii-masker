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

// validateStringMask reports whether result is all '*' characters with the same length as original
// (capped at maxLen if provided, otherwise just checks length equality).
func validateStringMask(result, original string, maxLen ...int) bool {
	cap := len(original)
	if len(maxLen) > 0 && maxLen[0] < cap {
		cap = maxLen[0]
	}
	if len(result) != cap {
		return false
	}
	for _, c := range result {
		if c != '*' {
			return false
		}
	}
	return true
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
