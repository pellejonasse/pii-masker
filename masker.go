package piimasker

import (
	"math/rand/v2"
	"reflect"
)

const (
	piiTagField = "Pii" // PiiMasking field tag, e.g. `Pii:"mask"` or `Pii:"anonymize"` not setting it will result in masking
)

type PiiMasker interface {
	Mask(obj any) any
}

// Don't necessarily need a struct and it might be more convenient if it isn't a struct
type piiMasker struct {
	config MaskerConfig
}

func NewMasker(config MaskerConfig) PiiMasker {
	return &piiMasker{config: config}
}

// generates a copy of an input object with all PII fields masked. The input can be a struct, map, slice, array, or any other type. The function uses reflection to traverse the input and apply masking to any fields that are tagged with `pii:"true"`. For strings, it replaces the content with a fixed mask (e.g., "****") while preserving the length of the original string up to a maximum defined length. For other types, it applies a generic mask (e.g., zero value for numbers, false for booleans). The function handles nested structures and collections recursively.
func (m *piiMasker) Mask(obj any) any {
	originalObject := reflect.ValueOf(obj)
	// Create a new value of the same type as the original
	copy := reflect.New(originalObject.Type()).Elem()
	m.recursiveStructTraverser(copy, originalObject, PiiModeNone)
	return copy.Interface()
}

// recursiveStructTraverser is a helper function that performs the actual traversal and masking of the struct fields. It takes three parameters: copy, which is the value being constructed as a copy of the original; original, which is the value being traversed; and piiValue, a boolean indicating whether the current field is tagged as PII. The function checks the type of each field and applies masking accordingly. For strings, it replaces the content with a fixed mask while preserving the length up to a maximum defined length. For other types, it applies a generic mask. The function handles nested structures and collections recursively, ensuring that all PII fields are masked throughout the entire object hierarchy.
func (m *piiMasker) recursiveStructTraverser(copy, original reflect.Value, piiMode PiiMode) {
	if original.CanInterface() {
		switch k := original.Kind(); {
		case k == reflect.String:
			applyStringPiiMode(m, copy, original, piiMode)
		case isInteger(k):
			applyIntPiiMode(copy, original, piiMode)
		case isFloat(k):
			applyFloatPiiMode(copy, original, piiMode)
		case isUnsignedInteger(k):
			applyUintPiiMode(copy, original, piiMode)
		case k == reflect.Bool:
			applyBoolPiiMode(copy, original, piiMode)
		case k == reflect.Slice || k == reflect.Array:
			if k == reflect.Slice {
				copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Len()))
			}
			for i := range original.Len() {
				m.recursiveStructTraverser(copy.Index(i), original.Index(i), piiMode)
			}

		case k == reflect.Pointer:
			pointerValue := original.Elem()
			if !pointerValue.IsValid() {
				// TODO: validate this
				return
			}
			copy.Set(reflect.New(pointerValue.Type()))
			m.recursiveStructTraverser(copy.Elem(), pointerValue, piiMode)

		case k == reflect.Map:
			// maps are a bit of a special case and we don't support setting pii on a map level
			copy.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				originalValue := original.MapIndex(key)
				copyValue := reflect.New(originalValue.Type()).Elem()
				m.recursiveStructTraverser(copyValue, originalValue, piiMode)
				copy.SetMapIndex(key, copyValue)
			}
		case k == reflect.Interface:
			originalValue := original.Elem()
			if !originalValue.IsValid() {
				copy.Set(reflect.Zero(original.Type()))
				return
			}
			copyValue := reflect.New(originalValue.Type()).Elem()
			m.recursiveStructTraverser(copyValue, originalValue, piiMode)
			copy.Set(copyValue)

		case k == reflect.Struct:
			for i := range original.NumField() {
				fieldTag := determinePiiMode(original.Type().Field(i).Tag.Get(piiTagField))
				if fieldTag == PiiModeNone {
					// no tag on this field — propagate parent mode
					m.recursiveStructTraverser(copy.Field(i), original.Field(i), piiMode)
				} else {
					// explicit tag — field tag wins
					m.recursiveStructTraverser(copy.Field(i), original.Field(i), fieldTag)
				}
			}
		}
	}
}

func isInteger(k reflect.Kind) bool {
	return k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64
}

func isFloat(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

func isUnsignedInteger(k reflect.Kind) bool {
	return k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64
}

func determinePiiMode(tag string) PiiMode {
	return PiiMode(tag)
}

func (m *piiMasker) maskString(s string) string {
	length := min(len(s), m.config.maxPiiStringLength)
	maskedString := make([]byte, length)

	for i := range maskedString {
		maskedString[i] = '*'
	}

	return string(maskedString)
}

func applyStringPiiMode(m *piiMasker, copy, original reflect.Value, piiMode PiiMode) {
	if original.Type().Name() != original.Kind().String() {
		copy.Set(reflect.ValueOf(original.Interface()))
		return
	}
	s := original.Interface().(string)
	switch piiMode {
	case PiiModeShow:
		copy.SetString(s)
	case PiiModeAnonymize:
		copy.SetString(anonymizeString(s))
	default:
		copy.SetString(m.maskString(s))
	}
}

func applyIntPiiMode(copy, original reflect.Value, piiMode PiiMode) {
	switch piiMode {
	case PiiModeShow:
		copy.SetInt(original.Int())
	case PiiModeAnonymize:
		copy.SetInt(anonymizeInt(original.Int()))
	default:
		copy.SetInt(0)
	}
}

func applyUintPiiMode(copy, original reflect.Value, piiMode PiiMode) {
	switch piiMode {
	case PiiModeShow:
		copy.SetUint(original.Uint())
	case PiiModeAnonymize:
		copy.SetUint(anonymizeUint(original.Uint()))
	default:
		copy.SetUint(0)
	}
}

func applyFloatPiiMode(copy, original reflect.Value, piiMode PiiMode) {
	switch piiMode {
	case PiiModeShow:
		copy.SetFloat(original.Float())
	case PiiModeAnonymize:
		copy.SetFloat(anonymizeFloat(original.Float()))
	default:
		copy.SetFloat(0)
	}
}

// bool is a bit of a special case since there are only 2 values, so its probably unlikely that this field will be PII, so we just default to false if it is a PII
func applyBoolPiiMode(copy, original reflect.Value, piiMode PiiMode) {
	switch piiMode {
	case PiiModeShow:
		copy.SetBool(original.Bool())
	case PiiModeAnonymize:
		copy.SetBool(rand.IntN(2) == 1)
	default:
		copy.SetBool(false)
	}
}
