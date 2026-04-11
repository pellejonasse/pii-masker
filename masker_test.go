package piimasker_test

import (
	"testing"
)

func TestMask_DefaultBehaviour(t *testing.T) {
	masker := newTestMasker(t)

	t.Run("string", func(t *testing.T) {
		type input struct{ Value string }
		original := input{Value: "John Smith"}
		result := masker.Mask(original).(input)
		if result.Value != "" && result.Value == original.Value {
			t.Errorf("expected Value to be masked, got %q", result.Value)
		}
	})

	t.Run("int", func(t *testing.T) {
		type input struct{ Value int }
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value != 0 {
			t.Errorf("expected Value to be masked to 0, got %d", result.Value)
		}
	})

	t.Run("uint", func(t *testing.T) {
		type input struct{ Value uint }
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value != 0 {
			t.Errorf("expected Value to be masked to 0, got %d", result.Value)
		}
	})

	t.Run("float", func(t *testing.T) {
		type input struct{ Value float64 }
		original := input{Value: 3.14}
		result := masker.Mask(original).(input)
		if result.Value != 0 {
			t.Errorf("expected Value to be masked to 0, got %f", result.Value)
		}
	})

	t.Run("bool", func(t *testing.T) {
		type input struct{ Value bool }
		original := input{Value: true}
		result := masker.Mask(original).(input)
		if result.Value != false {
			t.Errorf("expected Value to be masked to false, got %v", result.Value)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		type input struct{ Value *string }
		s := "secret"
		original := input{Value: &s}
		result := masker.Mask(original).(input)
		if result.Value == nil {
			t.Fatal("expected pointer to be non-nil")
		}
		if *result.Value == *original.Value {
			t.Errorf("expected pointed value to be masked, got %q", *result.Value)
		}
	})

	t.Run("slice", func(t *testing.T) {
		type input struct{ Value []string }
		original := input{Value: []string{"a", "b"}}
		result := masker.Mask(original).(input)
		if len(result.Value) != len(original.Value) {
			t.Errorf("expected slice length %d, got %d", len(original.Value), len(result.Value))
		}
	})

	t.Run("map", func(t *testing.T) {
		type input struct{ Value map[string]string }
		original := input{Value: map[string]string{"key": "value"}}
		result := masker.Mask(original).(input)
		if len(result.Value) != len(original.Value) {
			t.Errorf("expected map length %d, got %d", len(original.Value), len(result.Value))
		}
	})

	t.Run("nested", func(t *testing.T) {
		type inner struct{ Value string }
		type input struct{ Inner inner }
		original := input{Inner: inner{Value: "secret"}}
		result := masker.Mask(original).(input)
		if result.Inner.Value == original.Inner.Value {
			t.Errorf("expected nested Value to be masked, got %q", result.Inner.Value)
		}
	})
}

func TestMask_MaskTag(t *testing.T) {
	masker := newTestMasker(t)

	t.Run("string", func(t *testing.T) {
		type input struct {
			Value string `Pii:"mask"`
		}
		original := input{Value: "John Smith"}
		result := masker.Mask(original).(input)
		if result.Value == original.Value {
			t.Errorf("expected Value to be masked, got %q", result.Value)
		}
	})

	t.Run("int", func(t *testing.T) {
		type input struct {
			Value int `Pii:"mask"`
		}
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value != 0 {
			t.Errorf("expected Value to be masked to 0, got %d", result.Value)
		}
	})

	t.Run("uint", func(t *testing.T) {
		type input struct {
			Value uint `Pii:"mask"`
		}
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value != 0 {
			t.Errorf("expected Value to be masked to 0, got %d", result.Value)
		}
	})

	t.Run("float", func(t *testing.T) {
		type input struct {
			Value float64 `Pii:"mask"`
		}
		original := input{Value: 3.14}
		result := masker.Mask(original).(input)
		if result.Value != 0 {
			t.Errorf("expected Value to be masked to 0, got %f", result.Value)
		}
	})

	t.Run("bool", func(t *testing.T) {
		type input struct {
			Value bool `Pii:"mask"`
		}
		original := input{Value: true}
		result := masker.Mask(original).(input)
		if result.Value != false {
			t.Errorf("expected Value to be masked to false, got %v", result.Value)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		type input struct {
			Value *string `Pii:"mask"`
		}
		s := "secret"
		original := input{Value: &s}
		result := masker.Mask(original).(input)
		if result.Value == nil {
			t.Fatal("expected pointer to be non-nil")
		}
		if *result.Value == *original.Value {
			t.Errorf("expected pointed value to be masked, got %q", *result.Value)
		}
	})

	t.Run("nested", func(t *testing.T) {
		type inner struct {
			Value string `Pii:"mask"`
		}
		type input struct{ Inner inner }
		original := input{Inner: inner{Value: "secret"}}
		result := masker.Mask(original).(input)
		if result.Inner.Value == original.Inner.Value {
			t.Errorf("expected nested Value to be masked, got %q", result.Inner.Value)
		}
	})
}

func TestMask_AnonymizeTag(t *testing.T) {
	masker := newTestMasker(t)

	t.Run("string", func(t *testing.T) {
		type input struct {
			Value string `Pii:"anonymize"`
		}
		original := input{Value: "John Smith"}
		result := masker.Mask(original).(input)
		if result.Value == original.Value {
			t.Errorf("expected Value to be anonymized, got same value %q", result.Value)
		}
		if len(result.Value) != len(original.Value) {
			t.Errorf("expected anonymized length %d, got %d", len(original.Value), len(result.Value))
		}
	})

	t.Run("int", func(t *testing.T) {
		type input struct {
			Value int `Pii:"anonymize"`
		}
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value == 0 {
			t.Error("expected anonymized int to be non-zero")
		}
	})

	t.Run("uint", func(t *testing.T) {
		type input struct {
			Value uint `Pii:"anonymize"`
		}
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value == 0 {
			t.Error("expected anonymized uint to be non-zero")
		}
	})

	t.Run("float", func(t *testing.T) {
		type input struct {
			Value float64 `Pii:"anonymize"`
		}
		original := input{Value: 3.14}
		result := masker.Mask(original).(input)
		if result.Value == 0 {
			t.Error("expected anonymized float to be non-zero")
		}
	})

	t.Run("bool", func(t *testing.T) {
		type input struct {
			Value bool `Pii:"anonymize"`
		}
		// run several times since random bool could coincidentally match
		masker := newTestMasker(t)
		allSame := true
		for range 20 {
			original := input{Value: true}
			result := masker.Mask(original).(input)
			if result.Value != original.Value {
				allSame = false
				break
			}
		}
		if allSame {
			t.Error("expected anonymized bool to differ from original at least once in 20 runs")
		}
	})

	t.Run("pointer", func(t *testing.T) {
		type input struct {
			Value *string `Pii:"anonymize"`
		}
		s := "secret"
		original := input{Value: &s}
		result := masker.Mask(original).(input)
		if result.Value == nil {
			t.Fatal("expected pointer to be non-nil")
		}
		if *result.Value == *original.Value {
			t.Errorf("expected pointed value to be anonymized, got %q", *result.Value)
		}
	})

	t.Run("nested", func(t *testing.T) {
		type inner struct {
			Value string `Pii:"anonymize"`
		}
		type input struct{ Inner inner }
		original := input{Inner: inner{Value: "secret"}}
		result := masker.Mask(original).(input)
		if result.Inner.Value == original.Inner.Value {
			t.Errorf("expected nested Value to be anonymized, got %q", result.Inner.Value)
		}
	})
}

func TestMask_ShowTag(t *testing.T) {
	masker := newTestMasker(t)

	t.Run("string", func(t *testing.T) {
		type input struct {
			Value string `Pii:"show"`
		}
		original := input{Value: "John Smith"}
		result := masker.Mask(original).(input)
		if result.Value != original.Value {
			t.Errorf("expected Value to be preserved, got %q", result.Value)
		}
	})

	t.Run("int", func(t *testing.T) {
		type input struct {
			Value int `Pii:"show"`
		}
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value != original.Value {
			t.Errorf("expected Value to be preserved, got %d", result.Value)
		}
	})

	t.Run("uint", func(t *testing.T) {
		type input struct {
			Value uint `Pii:"show"`
		}
		original := input{Value: 42}
		result := masker.Mask(original).(input)
		if result.Value != original.Value {
			t.Errorf("expected Value to be preserved, got %d", result.Value)
		}
	})

	t.Run("float", func(t *testing.T) {
		type input struct {
			Value float64 `Pii:"show"`
		}
		original := input{Value: 3.14}
		result := masker.Mask(original).(input)
		if result.Value != original.Value {
			t.Errorf("expected Value to be preserved, got %f", result.Value)
		}
	})

	t.Run("bool", func(t *testing.T) {
		type input struct {
			Value bool `Pii:"show"`
		}
		original := input{Value: true}
		result := masker.Mask(original).(input)
		if result.Value != original.Value {
			t.Errorf("expected Value to be preserved, got %v", result.Value)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		type input struct {
			Value *string `Pii:"show"`
		}
		s := "secret"
		original := input{Value: &s}
		result := masker.Mask(original).(input)
		if result.Value == nil {
			t.Fatal("expected pointer to be non-nil")
		}
		if *result.Value != *original.Value {
			t.Errorf("expected pointed value to be preserved, got %q", *result.Value)
		}
	})

	t.Run("nested", func(t *testing.T) {
		type inner struct {
			Value string `Pii:"show"`
		}
		type input struct{ Inner inner }
		original := input{Inner: inner{Value: "secret"}}
		result := masker.Mask(original).(input)
		if result.Inner.Value != original.Inner.Value {
			t.Errorf("expected nested Value to be preserved, got %q", result.Inner.Value)
		}
	})
}
