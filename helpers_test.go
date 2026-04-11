package piimasker_test

import (
	"testing"

	piimasker "piimasker"
)

// could potentially pass this as a parameter so you can test multiple settings, in fact
var testConfig = piimasker.MaskerConfig{MaxPiiStringLength: 100}

// newTestMasker creates a Masker with default config for use in tests and benchmarks.
func newTestMasker(tb testing.TB, config ...piimasker.MaskerConfig) piimasker.PiiMasker {
	tb.Helper()
	if len(config) == 0 {
		return piimasker.NewMasker(testConfig)
	}
	return piimasker.NewMasker(config[0])
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

// validateAnonymization reports whether result has the same length as original (capped at maxLen
// if provided) and has different content.
func validateAnonymization(result, original string, maxLen ...int) bool {
	cap := len(original)
	if len(maxLen) > 0 && maxLen[0] < cap {
		cap = maxLen[0]
	}
	return len(result) == cap && result != original
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
