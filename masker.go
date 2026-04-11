package piimasker

import (
	"math/rand/v2"
	"reflect"
	"sync"
)

const (
	piiTagField = "Pii" // PiiMasking field tag, e.g. `Pii:"mask"` or `Pii:"anonymize"` not setting it will result in masking
)

type PiiMasker interface {
	Mask(obj any, opts ...MaskerConfig) any
}

// Don't necessarily need a struct and it might be more convenient if it isn't a struct, since now I have to worry about global state
type piiMasker struct {
	config    MaskerConfig
	typeCache sync.Map // map[reflect.Type][]piiMode, populated on first encounter of each struct type
}

func NewMasker(config MaskerConfig) PiiMasker {
	if config.MaxPiiStringLength == 0 {
		config.MaxPiiStringLength = 100
	}
	return &piiMasker{config: config}
}

// generates a copy of an input object with all PII fields masked. The input can be a struct, map, slice, array, or any other type. The function uses reflection to traverse the input and apply masking to any fields that are tagged with `pii:"true"`. For strings, it replaces the content with a fixed mask (e.g., "****") while preserving the length of the original string up to a maximum defined length. For other types, it applies a generic mask (e.g., zero value for numbers, false for booleans). The function handles nested structures and collections recursively.
func (m *piiMasker) Mask(obj any, opts ...MaskerConfig) any {
	config := m.config
	if len(opts) > 0 {
		config = opts[0]
		if config.MaxPiiStringLength == 0 {
			config.MaxPiiStringLength = m.config.MaxPiiStringLength
		}
	}
	originalObject := reflect.ValueOf(obj)
	copy := reflect.New(originalObject.Type()).Elem()
	m.recursiveStructTraverserWithConfig(copy, originalObject, piiModeNone, config)
	return copy.Interface()
}

// recursiveStructTraverser is a helper function that performs the actual traversal and masking of the struct fields. It takes three parameters: copy, which is the value being constructed as a copy of the original; original, which is the value being traversed; and piiValue, a boolean indicating whether the current field is tagged as PII. The function checks the type of each field and applies masking accordingly. For strings, it replaces the content with a fixed mask while preserving the length up to a maximum defined length. For other types, it applies a generic mask. The function handles nested structures and collections recursively, ensuring that all PII fields are masked throughout the entire object hierarchy.
// NOTE: Should I add support for time.Time and raw json strings?
func (m *piiMasker) recursiveStructTraverserWithConfig(copy, original reflect.Value, piiMode piiMode, config MaskerConfig) {
	if original.CanInterface() {
		switch k := original.Kind(); {
		case k == reflect.String:
			applyStringPiiMode(copy, original, piiMode, config)
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
				m.recursiveStructTraverserWithConfig(copy.Index(i), original.Index(i), piiMode, config)
			}

		case k == reflect.Pointer:
			pointerValue := original.Elem()
			if !pointerValue.IsValid() {
				return
			}
			copy.Set(reflect.New(pointerValue.Type()))
			m.recursiveStructTraverserWithConfig(copy.Elem(), pointerValue, piiMode, config)

		case k == reflect.Map:
			copy.Set(reflect.MakeMap(original.Type()))
			for _, key := range original.MapKeys() {
				originalValue := original.MapIndex(key)
				copyValue := reflect.New(originalValue.Type()).Elem()
				m.recursiveStructTraverserWithConfig(copyValue, originalValue, piiMode, config)
				copy.SetMapIndex(key, copyValue)
			}
		case k == reflect.Interface:
			originalValue := original.Elem()
			if !originalValue.IsValid() {
				copy.Set(reflect.Zero(original.Type()))
				return
			}
			copyValue := reflect.New(originalValue.Type()).Elem()
			m.recursiveStructTraverserWithConfig(copyValue, originalValue, piiMode, config)
			copy.Set(copyValue)

		case k == reflect.Struct:
			fieldTags := m.structFieldTags(original.Type())
			for i := range original.NumField() {
				if fieldTags[i] == piiModeNone {
					m.recursiveStructTraverserWithConfig(copy.Field(i), original.Field(i), piiMode, config)
				} else {
					m.recursiveStructTraverserWithConfig(copy.Field(i), original.Field(i), fieldTags[i], config)
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

func determinePiiMode(tag string) piiMode {
	return piiMode(tag)
}

func (m *piiMasker) structFieldTags(t reflect.Type) []piiMode {
	if cached, ok := m.typeCache.Load(t); ok {
		return cached.([]piiMode)
	}
	tags := make([]piiMode, t.NumField())
	for i := range t.NumField() {
		tags[i] = determinePiiMode(t.Field(i).Tag.Get(piiTagField))
	}
	actual, _ := m.typeCache.LoadOrStore(t, tags)
	return actual.([]piiMode)
}

func maskString(s string, config MaskerConfig) string {
	length := min(len(s), config.MaxPiiStringLength)
	maskedString := make([]byte, length)
	for i := range maskedString {
		maskedString[i] = '*'
	}
	return string(maskedString)
}

func applyStringPiiMode(copy, original reflect.Value, piiMode piiMode, config MaskerConfig) {
	s := original.String()
	switch piiMode {
	case piiModeShow:
		copy.SetString(s)
	case piiModeAnonymize:
		copy.SetString(anonymizeString(s, config))
	default:
		copy.SetString(maskString(s, config))
	}
}

func applyIntPiiMode(copy, original reflect.Value, piiMode piiMode) {
	switch piiMode {
	case piiModeShow:
		copy.SetInt(original.Int())
	case piiModeAnonymize:
		copy.SetInt(anonymizeInt(original.Int()))
	default:
		copy.SetInt(0)
	}
}

func applyUintPiiMode(copy, original reflect.Value, piiMode piiMode) {
	switch piiMode {
	case piiModeShow:
		copy.SetUint(original.Uint())
	case piiModeAnonymize:
		copy.SetUint(anonymizeUint(original.Uint()))
	default:
		copy.SetUint(0)
	}
}

func applyFloatPiiMode(copy, original reflect.Value, piiMode piiMode) {
	switch piiMode {
	case piiModeShow:
		copy.SetFloat(original.Float())
	case piiModeAnonymize:
		copy.SetFloat(anonymizeFloat(original.Float()))
	default:
		copy.SetFloat(0)
	}
}

// bool is a bit of a special case since there are only 2 values, so its probably unlikely that this field will be PII, so we just default to false if it is a PII
func applyBoolPiiMode(copy, original reflect.Value, piiMode piiMode) {
	switch piiMode {
	case piiModeShow:
		copy.SetBool(original.Bool())
	case piiModeAnonymize:
		copy.SetBool(rand.IntN(2) == 1)
	default:
		copy.SetBool(false)
	}
}
