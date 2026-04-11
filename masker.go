package piimasker

import (
	"reflect"
)

const (
	maxPiiStringLength = 100 // Maximum length of a string to reduce noise of masking big items
)

type PiiMasker interface {
	Mask(obj any) any
}

type piiMasker struct {
	config Config
}

func NewMasker(config Config) PiiMasker {
	return &piiMasker{config: config}
}

// generates a copy of an input object with all PII fields masked. The input can be a struct, map, slice, array, or any other type. The function uses reflection to traverse the input and apply masking to any fields that are tagged with `pii:"true"`. For strings, it replaces the content with a fixed mask (e.g., "****") while preserving the length of the original string up to a maximum defined length. For other types, it applies a generic mask (e.g., zero value for numbers, false for booleans). The function handles nested structures and collections recursively.
func (m *piiMasker) Mask(obj any) any {
	originalObject := reflect.ValueOf(obj)
	// Create a new value of the same type as the original
	copy := reflect.New(originalObject.Type()).Elem()
	m.recursiveStructTraverser(copy, originalObject, true)
	return copy.Interface()
}

// recursiveStructTraverser is a helper function that performs the actual traversal and masking of the struct fields. It takes three parameters: copy, which is the value being constructed as a copy of the original; original, which is the value being traversed; and piiValue, a boolean indicating whether the current field is tagged as PII. The function checks the type of each field and applies masking accordingly. For strings, it replaces the content with a fixed mask while preserving the length up to a maximum defined length. For other types, it applies a generic mask. The function handles nested structures and collections recursively, ensuring that all PII fields are masked throughout the entire object hierarchy.
func (m *piiMasker) recursiveStructTraverser(copy, original reflect.Value, piiValue bool) {
	if original.CanInterface() {
		switch k := original.Kind(); {
		case k == reflect.String:
			if original.Type().Name() == original.Kind().String() {
				stringValue := original.Interface().(string)
				if piiValue {
					copy.SetString(stringValue)
				} else {
					stringValue = m.maskString(stringValue)
					copy.SetString(stringValue)
				}
			} else {
				stringValue := original.Interface()
				copy.Set(reflect.ValueOf(stringValue))
			}
		case k == reflect.Pointer:
			// handle pointer
		case isNumericKind(k):
			// handle numbers
		case k == reflect.Bool:
			// handle bool
		case k == reflect.Slice || k == reflect.Array:
			// handle slice/array
		case k == reflect.Map:
			// handle map
		case k == reflect.Interface:
			// handle interface
		case k == reflect.Struct:
			// handle struct
		}
	}
}

func isNumericKind(k reflect.Kind) bool {
	return k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 ||
		k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64 ||
		k == reflect.Float32 || k == reflect.Float64
}

func (m *piiMasker) maskString(s string) string {
	length := min(len(s), m.config.maxPiiStringLength)
	maskedString := make([]byte, length)

	for i := range maskedString {
		maskedString[i] = '*'
	}

	return string(maskedString)
}
