package piimasker

import (
	"reflect"
	"testing"
)

type cacheStructA struct {
	Name  string `Pii:"mask"`
	Email string `Pii:"mask"`
	Age   int    `Pii:"show"`
}

type cacheStructB struct {
	Token   string  `Pii:"anonymize"`
	Balance float64 `Pii:"show"`
	Active  bool
}

type cacheStructC struct {
	DeviceID string `Pii:"mask"`
	IP       string `Pii:"anonymize"`
	Score    float64
}

func TestTypeCache_Contents(t *testing.T) {
	m := NewMasker(MaskerConfig{}).(*piiMasker)

	m.Mask(cacheStructA{Name: "Alice", Email: "alice@example.com", Age: 30})
	m.Mask(cacheStructB{Token: "tok", Balance: 100.0, Active: true})
	m.Mask(cacheStructC{DeviceID: "dev", IP: "1.2.3.4", Score: 9.5})

	snapshot := map[string]map[string]piiMode{}
	m.typeCache.Range(func(key, value any) bool {
		rt := key.(reflect.Type)
		tags := value.([]piiMode)
		fields := map[string]piiMode{}
		for i, tag := range tags {
			mode := tag
			if mode == piiModeNone {
				mode = "(none)"
			}
			fields[rt.Field(i).Name] = mode
		}
		snapshot[rt.Name()] = fields
		return true
	})
	t.Logf("typeCache: %+v", snapshot)

	count := 0
	m.typeCache.Range(func(_, _ any) bool { count++; return true })
	if count != 3 {
		t.Errorf("expected 3 entries in typeCache, got %d", count)
	}
}
