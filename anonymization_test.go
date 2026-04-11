package piimasker_test

import (
	"fmt"
	"testing"
)

func TestAnonymize_Correctness(t *testing.T) {
	masker := newTestMasker(t)

	t.Run("string", func(t *testing.T) {
		type input struct {
			Name string `Pii:"anonymize"`
		}
		original := input{Name: "John Smith"}
		result := masker.Mask(original).(input)
		if result.Name == original.Name {
			t.Errorf("expected anonymized value to differ from %q", original.Name)
		}
		if len(result.Name) != len(original.Name) {
			t.Errorf("expected length %d, got %d", len(original.Name), len(result.Name))
		}
	})

	t.Run("int", func(t *testing.T) {
		type input struct {
			Age int `Pii:"anonymize"`
		}
		original := input{Age: 12345}
		result := masker.Mask(original).(input)
		if result.Age == 0 {
			t.Error("expected anonymized int to be non-zero")
		}
		if len(fmt.Sprintf("%d", result.Age)) != len("12345") {
			t.Errorf("expected 5 digits, got %d", len(fmt.Sprintf("%d", result.Age)))
		}
	})

	t.Run("uint", func(t *testing.T) {
		type input struct {
			ID uint `Pii:"anonymize"`
		}
		original := input{ID: 98765}
		result := masker.Mask(original).(input)
		if result.ID == 0 {
			t.Error("expected anonymized uint to be non-zero")
		}
		if len(fmt.Sprintf("%d", result.ID)) != len("98765") {
			t.Errorf("expected 5 digits, got %d", len(fmt.Sprintf("%d", result.ID)))
		}
	})

	t.Run("float", func(t *testing.T) {
		type input struct {
			Score float64 `Pii:"anonymize"`
		}
		original := input{Score: 123.45}
		result := masker.Mask(original).(input)
		if result.Score == 0 {
			t.Error("expected anonymized float to be non-zero")
		}
	})
}

func BenchmarkAnonymize(b *testing.B) {
	masker := newTestMasker(b)

	b.Run("string", func(b *testing.B) {
		type input struct {
			Name string `Pii:"anonymize"`
		}
		original := input{Name: "John Smith"}
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(original)
		}
	})

	b.Run("int", func(b *testing.B) {
		type input struct {
			Age int `Pii:"anonymize"`
		}
		original := input{Age: 12345}
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(original)
		}
	})

	b.Run("uint", func(b *testing.B) {
		type input struct {
			ID uint `Pii:"anonymize"`
		}
		original := input{ID: 98765}
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(original)
		}
	})

	b.Run("float", func(b *testing.B) {
		type input struct {
			Score float64 `Pii:"anonymize"`
		}
		original := input{Score: 123.45}
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(original)
		}
	})
}
