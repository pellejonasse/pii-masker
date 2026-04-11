package piimasker_test

import (
	"fmt"
	"piimasker"
	"testing"
)

func TestMask_DefaultBehaviour(t *testing.T) {
	configs := []piimasker.MaskerConfig{
		{MaxPiiStringLength: 5},
		{MaxPiiStringLength: 10},
		{MaxPiiStringLength: 20},
		{MaxPiiStringLength: 100},
	}
	for _, cfg := range configs {
		cfg := cfg
		t.Run(fmt.Sprintf("maxLen=%d", cfg.MaxPiiStringLength), func(t *testing.T) {
			t.Parallel()
			masker := newTestMasker(t, cfg)

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
		})
	}
}

func TestMask_MaskTag(t *testing.T) {
	configs := []piimasker.MaskerConfig{
		{MaxPiiStringLength: 5},
		{MaxPiiStringLength: 10},
		{MaxPiiStringLength: 20},
		{MaxPiiStringLength: 100},
	}
	for _, cfg := range configs {
		cfg := cfg
		t.Run(fmt.Sprintf("maxLen=%d", cfg.MaxPiiStringLength), func(t *testing.T) {
			t.Parallel()
			masker := newTestMasker(t, cfg)

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
		})
	}
}

func TestMask_AnonymizeTag(t *testing.T) {
	configs := []piimasker.MaskerConfig{
		{MaxPiiStringLength: 5},
		{MaxPiiStringLength: 10},
		{MaxPiiStringLength: 20},
		{MaxPiiStringLength: 100},
	}
	for _, cfg := range configs {
		cfg := cfg
		t.Run(fmt.Sprintf("maxLen=%d", cfg.MaxPiiStringLength), func(t *testing.T) {
			t.Parallel()
			masker := newTestMasker(t, cfg)

			t.Run("string", func(t *testing.T) {
				type input struct {
					Value string `Pii:"anonymize"`
				}
				original := input{Value: "John Smith"}
				result := masker.Mask(original).(input)
				if !validateAnonymization(result.Value, original.Value, cfg.MaxPiiStringLength) {
					t.Errorf("expected Value to be anonymized, got %q", result.Value)
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
		})
	}
}

func TestMask_ShowTag(t *testing.T) {
	configs := []piimasker.MaskerConfig{
		{MaxPiiStringLength: 5},
		{MaxPiiStringLength: 10},
		{MaxPiiStringLength: 20},
		{MaxPiiStringLength: 100},
	}
	for _, cfg := range configs {
		cfg := cfg
		t.Run(fmt.Sprintf("maxLen=%d", cfg.MaxPiiStringLength), func(t *testing.T) {
			t.Parallel()
			masker := newTestMasker(t, cfg)

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
		})
	}
}

func TestMask_Integration_MaskTag(t *testing.T) {
	m := newTestMasker(t)
	result := m.Mask(newPersonFixture()).(Person)

	t.Run("top_level", func(t *testing.T) {
		if !validateStringMask(result.FirstName, fixtureFirstName, testConfig.MaxPiiStringLength) {
			t.Errorf("FirstName: expected masked, got %q", result.FirstName)
		}
		if !validateStringMask(result.LastName, fixtureLastName, testConfig.MaxPiiStringLength) {
			t.Errorf("LastName: expected masked, got %q", result.LastName)
		}
		if result.Age != 0 {
			t.Errorf("Age: expected masked to 0, got %d", result.Age)
		}
		if result.IsActive != false {
			t.Errorf("IsActive: expected masked to false, got %v", result.IsActive)
		}
	})

	t.Run("contact", func(t *testing.T) {
		c := result.Contact
		if !validateStringMask(c.Email, fixtureEmail, testConfig.MaxPiiStringLength) {
			t.Errorf("Email: expected masked, got %q", c.Email)
		}
		if !validateStringMask(c.Phone, fixturePhone, testConfig.MaxPiiStringLength) {
			t.Errorf("Phone: expected masked, got %q", c.Phone)
		}
		if !validateStringMask(c.AltEmail, fixtureAltEmail, testConfig.MaxPiiStringLength) {
			t.Errorf("AltEmail: expected masked, got %q", c.AltEmail)
		}
	})

	t.Run("address", func(t *testing.T) {
		a := result.Contact.Address
		if !validateStringMask(a.Street, fixtureStreet, testConfig.MaxPiiStringLength) {
			t.Errorf("Street: expected masked, got %q", a.Street)
		}
		if !validateStringMask(a.City, fixtureCity, testConfig.MaxPiiStringLength) {
			t.Errorf("City: expected masked, got %q", a.City)
		}
		if !validateStringMask(a.Country, fixtureCountry, testConfig.MaxPiiStringLength) {
			t.Errorf("Country: expected masked, got %q", a.Country)
		}
		if !validateStringMask(a.ZipCode, fixtureZipCode, testConfig.MaxPiiStringLength) {
			t.Errorf("ZipCode: expected masked, got %q", a.ZipCode)
		}
	})

	t.Run("coordinates", func(t *testing.T) {
		c := result.Contact.Address.Coordinates
		if c.Latitude != 0 {
			t.Errorf("Latitude: expected masked to 0, got %f", c.Latitude)
		}
		if c.Longitude != 0 {
			t.Errorf("Longitude: expected masked to 0, got %f", c.Longitude)
		}
		if c.Altitude != 0 {
			t.Errorf("Altitude: expected masked to 0, got %f", c.Altitude)
		}
		if c.Accuracy != 0 {
			t.Errorf("Accuracy: expected masked to 0, got %f", c.Accuracy)
		}
	})

	t.Run("payment_method", func(t *testing.T) {
		pm := result.PaymentMethods[0]
		if !validateStringMask(pm.CardNumber, fixtureCardNumber, testConfig.MaxPiiStringLength) {
			t.Errorf("CardNumber: expected masked, got %q", pm.CardNumber)
		}
		if !validateStringMask(pm.CVV, fixtureCVV, testConfig.MaxPiiStringLength) {
			t.Errorf("CVV: expected masked, got %q", pm.CVV)
		}
		if !validateStringMask(pm.Expiry, fixtureExpiry, testConfig.MaxPiiStringLength) {
			t.Errorf("Expiry: expected masked, got %q", pm.Expiry)
		}
		if !validateStringMask(pm.HolderName, fixtureHolderName, testConfig.MaxPiiStringLength) {
			t.Errorf("HolderName: expected masked, got %q", pm.HolderName)
		}
		if pm.IsDefault != false {
			t.Errorf("IsDefault: expected masked to false, got %v", pm.IsDefault)
		}
		ba := pm.BillingAddress
		if !validateStringMask(ba.Street, fixtureStreet, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.Street: expected masked, got %q", ba.Street)
		}
		if !validateStringMask(ba.City, fixtureCity, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.City: expected masked, got %q", ba.City)
		}
		if !validateStringMask(ba.PostCode, fixtureZipCode, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.PostCode: expected masked, got %q", ba.PostCode)
		}
		if !validateStringMask(ba.Country, fixtureCountry, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.Country: expected masked, got %q", ba.Country)
		}
	})

	t.Run("order", func(t *testing.T) {
		o := result.Orders[0]
		if !validateStringMask(o.OrderID, fixtureOrderID, testConfig.MaxPiiStringLength) {
			t.Errorf("OrderID: expected masked, got %q", o.OrderID)
		}
		if o.Amount != 0 {
			t.Errorf("Amount: expected masked to 0, got %f", o.Amount)
		}
		if !validateStringMask(o.Currency, fixtureCurrency, testConfig.MaxPiiStringLength) {
			t.Errorf("Currency: expected masked, got %q", o.Currency)
		}
		if !validateStringMask(o.Notes, fixtureNotes, testConfig.MaxPiiStringLength) {
			t.Errorf("Notes: expected masked, got %q", o.Notes)
		}
		item := o.Items[0]
		if !validateStringMask(item.ProductID, fixtureProductID, testConfig.MaxPiiStringLength) {
			t.Errorf("ProductID: expected masked, got %q", item.ProductID)
		}
		if !validateStringMask(item.Name, fixtureItemName, testConfig.MaxPiiStringLength) {
			t.Errorf("Name: expected masked, got %q", item.Name)
		}
		if item.Quantity != 0 {
			t.Errorf("Quantity: expected masked to 0, got %d", item.Quantity)
		}
		if item.UnitPrice != 0 {
			t.Errorf("UnitPrice: expected masked to 0, got %f", item.UnitPrice)
		}
	})

	t.Run("device", func(t *testing.T) {
		d := result.Devices["mobile"]
		if !validateStringMask(d.DeviceID, fixtureDeviceID, testConfig.MaxPiiStringLength) {
			t.Errorf("DeviceID: expected masked, got %q", d.DeviceID)
		}
		if !validateStringMask(d.UserAgent, fixtureUserAgent, testConfig.MaxPiiStringLength) {
			t.Errorf("UserAgent: expected masked, got %q", d.UserAgent)
		}
		if !validateStringMask(d.IPAddress, fixtureIPAddress, testConfig.MaxPiiStringLength) {
			t.Errorf("IPAddress: expected masked, got %q", d.IPAddress)
		}
		if !validateStringMask(d.ScreenSize, fixtureScreenSize, testConfig.MaxPiiStringLength) {
			t.Errorf("ScreenSize: expected masked, got %q", d.ScreenSize)
		}
	})
}

func TestMask_Integration_ShowTag(t *testing.T) {
	m := newTestMasker(t)
	result := m.Mask(newPersonShowFixture()).(PersonShow)

	t.Run("top_level", func(t *testing.T) {
		if result.FirstName != fixtureFirstName {
			t.Errorf("FirstName: expected %q, got %q", fixtureFirstName, result.FirstName)
		}
		if result.LastName != fixtureLastName {
			t.Errorf("LastName: expected %q, got %q", fixtureLastName, result.LastName)
		}
		if result.Age != fixtureAge {
			t.Errorf("Age: expected %d, got %d", fixtureAge, result.Age)
		}
		if result.IsActive != true {
			t.Errorf("IsActive: expected true, got %v", result.IsActive)
		}
	})

	t.Run("contact", func(t *testing.T) {
		c := result.Contact
		if c.Email != fixtureEmail {
			t.Errorf("Email: expected %q, got %q", fixtureEmail, c.Email)
		}
		if c.Phone != fixturePhone {
			t.Errorf("Phone: expected %q, got %q", fixturePhone, c.Phone)
		}
		if c.AltEmail != fixtureAltEmail {
			t.Errorf("AltEmail: expected %q, got %q", fixtureAltEmail, c.AltEmail)
		}
	})

	t.Run("address", func(t *testing.T) {
		a := result.Contact.Address
		if a.Street != fixtureStreet {
			t.Errorf("Street: expected %q, got %q", fixtureStreet, a.Street)
		}
		if a.City != fixtureCity {
			t.Errorf("City: expected %q, got %q", fixtureCity, a.City)
		}
		if a.Country != fixtureCountry {
			t.Errorf("Country: expected %q, got %q", fixtureCountry, a.Country)
		}
		if a.ZipCode != fixtureZipCode {
			t.Errorf("ZipCode: expected %q, got %q", fixtureZipCode, a.ZipCode)
		}
	})

	t.Run("coordinates", func(t *testing.T) {
		c := result.Contact.Address.Coordinates
		if c.Latitude != fixtureLatitude {
			t.Errorf("Latitude: expected %f, got %f", fixtureLatitude, c.Latitude)
		}
		if c.Longitude != fixtureLongitude {
			t.Errorf("Longitude: expected %f, got %f", fixtureLongitude, c.Longitude)
		}
		if c.Altitude != fixtureAltitude {
			t.Errorf("Altitude: expected %f, got %f", fixtureAltitude, c.Altitude)
		}
		if c.Accuracy != fixtureAccuracy {
			t.Errorf("Accuracy: expected %f, got %f", fixtureAccuracy, c.Accuracy)
		}
	})

	t.Run("payment_method", func(t *testing.T) {
		pm := result.PaymentMethods[0]
		if pm.CardNumber != fixtureCardNumber {
			t.Errorf("CardNumber: expected %q, got %q", fixtureCardNumber, pm.CardNumber)
		}
		if pm.CVV != fixtureCVV {
			t.Errorf("CVV: expected %q, got %q", fixtureCVV, pm.CVV)
		}
		if pm.Expiry != fixtureExpiry {
			t.Errorf("Expiry: expected %q, got %q", fixtureExpiry, pm.Expiry)
		}
		if pm.HolderName != fixtureHolderName {
			t.Errorf("HolderName: expected %q, got %q", fixtureHolderName, pm.HolderName)
		}
		if pm.IsDefault != true {
			t.Errorf("IsDefault: expected true, got %v", pm.IsDefault)
		}
		ba := pm.BillingAddress
		if ba.Street != fixtureStreet {
			t.Errorf("BillingAddress.Street: expected %q, got %q", fixtureStreet, ba.Street)
		}
		if ba.City != fixtureCity {
			t.Errorf("BillingAddress.City: expected %q, got %q", fixtureCity, ba.City)
		}
		if ba.PostCode != fixtureZipCode {
			t.Errorf("BillingAddress.PostCode: expected %q, got %q", fixtureZipCode, ba.PostCode)
		}
		if ba.Country != fixtureCountry {
			t.Errorf("BillingAddress.Country: expected %q, got %q", fixtureCountry, ba.Country)
		}
	})

	t.Run("order", func(t *testing.T) {
		o := result.Orders[0]
		if o.OrderID != fixtureOrderID {
			t.Errorf("OrderID: expected %q, got %q", fixtureOrderID, o.OrderID)
		}
		if o.Amount != fixtureAmount {
			t.Errorf("Amount: expected %f, got %f", fixtureAmount, o.Amount)
		}
		if o.Currency != fixtureCurrency {
			t.Errorf("Currency: expected %q, got %q", fixtureCurrency, o.Currency)
		}
		if o.Notes != fixtureNotes {
			t.Errorf("Notes: expected %q, got %q", fixtureNotes, o.Notes)
		}
		item := o.Items[0]
		if item.ProductID != fixtureProductID {
			t.Errorf("ProductID: expected %q, got %q", fixtureProductID, item.ProductID)
		}
		if item.Name != fixtureItemName {
			t.Errorf("Name: expected %q, got %q", fixtureItemName, item.Name)
		}
		if item.Quantity != fixtureQuantity {
			t.Errorf("Quantity: expected %d, got %d", fixtureQuantity, item.Quantity)
		}
		if item.UnitPrice != fixtureUnitPrice {
			t.Errorf("UnitPrice: expected %f, got %f", fixtureUnitPrice, item.UnitPrice)
		}
	})

	t.Run("device", func(t *testing.T) {
		d := result.Devices["mobile"]
		if d.DeviceID != fixtureDeviceID {
			t.Errorf("DeviceID: expected %q, got %q", fixtureDeviceID, d.DeviceID)
		}
		if d.UserAgent != fixtureUserAgent {
			t.Errorf("UserAgent: expected %q, got %q", fixtureUserAgent, d.UserAgent)
		}
		if d.IPAddress != fixtureIPAddress {
			t.Errorf("IPAddress: expected %q, got %q", fixtureIPAddress, d.IPAddress)
		}
		if d.ScreenSize != fixtureScreenSize {
			t.Errorf("ScreenSize: expected %q, got %q", fixtureScreenSize, d.ScreenSize)
		}
	})
}

func TestMask_Integration_AnonymizeTag(t *testing.T) {
	m := newTestMasker(t)
	fixture := newPersonAnonymizeFixture()
	result := m.Mask(fixture).(PersonAnonymize)

	t.Run("top_level", func(t *testing.T) {
		if !validateAnonymization(result.FirstName, fixture.FirstName, testConfig.MaxPiiStringLength) {
			t.Errorf("FirstName: expected anonymized same-length string, got %q", result.FirstName)
		}
		if !validateAnonymization(result.LastName, fixture.LastName, testConfig.MaxPiiStringLength) {
			t.Errorf("LastName: expected anonymized same-length string, got %q", result.LastName)
		}
		if result.Age == 0 {
			t.Errorf("Age: expected anonymized non-zero, got 0")
		}
		// IsActive is a bool — skipped; non-deterministic single-run assertion, covered by TestMask_AnonymizeTag/bool
	})

	t.Run("contact", func(t *testing.T) {
		c := result.Contact
		if !validateAnonymization(c.Email, fixture.Contact.Email, testConfig.MaxPiiStringLength) {
			t.Errorf("Email: expected anonymized same-length string, got %q", c.Email)
		}
		if !validateAnonymization(c.Phone, fixture.Contact.Phone, testConfig.MaxPiiStringLength) {
			t.Errorf("Phone: expected anonymized same-length string, got %q", c.Phone)
		}
		if !validateAnonymization(c.AltEmail, fixture.Contact.AltEmail, testConfig.MaxPiiStringLength) {
			t.Errorf("AltEmail: expected anonymized same-length string, got %q", c.AltEmail)
		}
	})

	t.Run("address", func(t *testing.T) {
		a := result.Contact.Address
		fa := fixture.Contact.Address
		if !validateAnonymization(a.Street, fa.Street, testConfig.MaxPiiStringLength) {
			t.Errorf("Street: expected anonymized same-length string, got %q", a.Street)
		}
		if !validateAnonymization(a.City, fa.City, testConfig.MaxPiiStringLength) {
			t.Errorf("City: expected anonymized same-length string, got %q", a.City)
		}
		if !validateAnonymization(a.Country, fa.Country, testConfig.MaxPiiStringLength) {
			t.Errorf("Country: expected anonymized same-length string, got %q", a.Country)
		}
		if !validateAnonymization(a.ZipCode, fa.ZipCode, testConfig.MaxPiiStringLength) {
			t.Errorf("ZipCode: expected anonymized same-length string, got %q", a.ZipCode)
		}
	})

	t.Run("coordinates", func(t *testing.T) {
		c := result.Contact.Address.Coordinates
		if c.Latitude == 0 {
			t.Errorf("Latitude: expected anonymized non-zero, got 0")
		}
		if c.Longitude == 0 {
			t.Errorf("Longitude: expected anonymized non-zero, got 0")
		}
		if c.Altitude == 0 {
			t.Errorf("Altitude: expected anonymized non-zero, got 0")
		}
		if c.Accuracy == 0 {
			t.Errorf("Accuracy: expected anonymized non-zero, got 0")
		}
	})

	t.Run("payment_method", func(t *testing.T) {
		pm := result.PaymentMethods[0]
		fpm := fixture.PaymentMethods[0]
		if !validateAnonymization(pm.CardNumber, fpm.CardNumber, testConfig.MaxPiiStringLength) {
			t.Errorf("CardNumber: expected anonymized same-length string, got %q", pm.CardNumber)
		}
		if !validateAnonymization(pm.CVV, fpm.CVV, testConfig.MaxPiiStringLength) {
			t.Errorf("CVV: expected anonymized same-length string, got %q", pm.CVV)
		}
		if !validateAnonymization(pm.Expiry, fpm.Expiry, testConfig.MaxPiiStringLength) {
			t.Errorf("Expiry: expected anonymized same-length string, got %q", pm.Expiry)
		}
		if !validateAnonymization(pm.HolderName, fpm.HolderName, testConfig.MaxPiiStringLength) {
			t.Errorf("HolderName: expected anonymized same-length string, got %q", pm.HolderName)
		}
		// IsDefault is a bool — skipped; non-deterministic single-run assertion, covered by TestMask_AnonymizeTag/bool
		ba := pm.BillingAddress
		fba := fpm.BillingAddress
		if !validateAnonymization(ba.Street, fba.Street, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.Street: expected anonymized same-length string, got %q", ba.Street)
		}
		if !validateAnonymization(ba.City, fba.City, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.City: expected anonymized same-length string, got %q", ba.City)
		}
		if !validateAnonymization(ba.PostCode, fba.PostCode, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.PostCode: expected anonymized same-length string, got %q", ba.PostCode)
		}
		if !validateAnonymization(ba.Country, fba.Country, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.Country: expected anonymized same-length string, got %q", ba.Country)
		}
	})

	t.Run("order", func(t *testing.T) {
		o := result.Orders[0]
		fo := fixture.Orders[0]
		if !validateAnonymization(o.OrderID, fo.OrderID, testConfig.MaxPiiStringLength) {
			t.Errorf("OrderID: expected anonymized same-length string, got %q", o.OrderID)
		}
		if o.Amount == 0 {
			t.Errorf("Amount: expected anonymized non-zero, got 0")
		}
		if !validateAnonymization(o.Currency, fo.Currency, testConfig.MaxPiiStringLength) {
			t.Errorf("Currency: expected anonymized same-length string, got %q", o.Currency)
		}
		if !validateAnonymization(o.Notes, fo.Notes, testConfig.MaxPiiStringLength) {
			t.Errorf("Notes: expected anonymized same-length string, got %q", o.Notes)
		}
		item := o.Items[0]
		fitem := fo.Items[0]
		if !validateAnonymization(item.ProductID, fitem.ProductID, testConfig.MaxPiiStringLength) {
			t.Errorf("ProductID: expected anonymized same-length string, got %q", item.ProductID)
		}
		if !validateAnonymization(item.Name, fitem.Name, testConfig.MaxPiiStringLength) {
			t.Errorf("Name: expected anonymized same-length string, got %q", item.Name)
		}
		if item.Quantity == 0 {
			t.Errorf("Quantity: expected anonymized non-zero, got 0")
		}
		if item.UnitPrice == 0 {
			t.Errorf("UnitPrice: expected anonymized non-zero, got 0")
		}
	})

	t.Run("device", func(t *testing.T) {
		d := result.Devices["mobile"]
		fd := fixture.Devices["mobile"]
		if !validateAnonymization(d.DeviceID, fd.DeviceID, testConfig.MaxPiiStringLength) {
			t.Errorf("DeviceID: expected anonymized same-length string, got %q", d.DeviceID)
		}
		if !validateAnonymization(d.UserAgent, fd.UserAgent, testConfig.MaxPiiStringLength) {
			t.Errorf("UserAgent: expected anonymized same-length string, got %q", d.UserAgent)
		}
		if !validateAnonymization(d.IPAddress, fd.IPAddress, testConfig.MaxPiiStringLength) {
			t.Errorf("IPAddress: expected anonymized same-length string, got %q", d.IPAddress)
		}
		if !validateAnonymization(d.ScreenSize, fd.ScreenSize, testConfig.MaxPiiStringLength) {
			t.Errorf("ScreenSize: expected anonymized same-length string, got %q", d.ScreenSize)
		}
	})
}

func TestMask_Integration_NoTag(t *testing.T) {
	m := newTestMasker(t)
	result := m.Mask(newPersonNoTagFixture()).(PersonNoTag)

	t.Run("top_level", func(t *testing.T) {
		if !validateStringMask(result.FirstName, fixtureFirstName, testConfig.MaxPiiStringLength) {
			t.Errorf("FirstName: expected masked, got %q", result.FirstName)
		}
		if !validateStringMask(result.LastName, fixtureLastName, testConfig.MaxPiiStringLength) {
			t.Errorf("LastName: expected masked, got %q", result.LastName)
		}
		if result.Age != 0 {
			t.Errorf("Age: expected masked to 0, got %d", result.Age)
		}
		if result.IsActive != false {
			t.Errorf("IsActive: expected masked to false, got %v", result.IsActive)
		}
	})

	t.Run("contact", func(t *testing.T) {
		c := result.Contact
		if !validateStringMask(c.Email, fixtureEmail, testConfig.MaxPiiStringLength) {
			t.Errorf("Email: expected masked, got %q", c.Email)
		}
		if !validateStringMask(c.Phone, fixturePhone, testConfig.MaxPiiStringLength) {
			t.Errorf("Phone: expected masked, got %q", c.Phone)
		}
		if !validateStringMask(c.AltEmail, fixtureAltEmail, testConfig.MaxPiiStringLength) {
			t.Errorf("AltEmail: expected masked, got %q", c.AltEmail)
		}
	})

	t.Run("address", func(t *testing.T) {
		a := result.Contact.Address
		if !validateStringMask(a.Street, fixtureStreet, testConfig.MaxPiiStringLength) {
			t.Errorf("Street: expected masked, got %q", a.Street)
		}
		if !validateStringMask(a.City, fixtureCity, testConfig.MaxPiiStringLength) {
			t.Errorf("City: expected masked, got %q", a.City)
		}
		if !validateStringMask(a.Country, fixtureCountry, testConfig.MaxPiiStringLength) {
			t.Errorf("Country: expected masked, got %q", a.Country)
		}
		if !validateStringMask(a.ZipCode, fixtureZipCode, testConfig.MaxPiiStringLength) {
			t.Errorf("ZipCode: expected masked, got %q", a.ZipCode)
		}
	})

	t.Run("coordinates", func(t *testing.T) {
		c := result.Contact.Address.Coordinates
		if c.Latitude != 0 {
			t.Errorf("Latitude: expected masked to 0, got %f", c.Latitude)
		}
		if c.Longitude != 0 {
			t.Errorf("Longitude: expected masked to 0, got %f", c.Longitude)
		}
		if c.Altitude != 0 {
			t.Errorf("Altitude: expected masked to 0, got %f", c.Altitude)
		}
		if c.Accuracy != 0 {
			t.Errorf("Accuracy: expected masked to 0, got %f", c.Accuracy)
		}
	})

	t.Run("payment_method", func(t *testing.T) {
		pm := result.PaymentMethods[0]
		if !validateStringMask(pm.CardNumber, fixtureCardNumber, testConfig.MaxPiiStringLength) {
			t.Errorf("CardNumber: expected masked, got %q", pm.CardNumber)
		}
		if !validateStringMask(pm.CVV, fixtureCVV, testConfig.MaxPiiStringLength) {
			t.Errorf("CVV: expected masked, got %q", pm.CVV)
		}
		if !validateStringMask(pm.Expiry, fixtureExpiry, testConfig.MaxPiiStringLength) {
			t.Errorf("Expiry: expected masked, got %q", pm.Expiry)
		}
		if !validateStringMask(pm.HolderName, fixtureHolderName, testConfig.MaxPiiStringLength) {
			t.Errorf("HolderName: expected masked, got %q", pm.HolderName)
		}
		if pm.IsDefault != false {
			t.Errorf("IsDefault: expected masked to false, got %v", pm.IsDefault)
		}
		ba := pm.BillingAddress
		if !validateStringMask(ba.Street, fixtureStreet, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.Street: expected masked, got %q", ba.Street)
		}
		if !validateStringMask(ba.City, fixtureCity, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.City: expected masked, got %q", ba.City)
		}
		if !validateStringMask(ba.PostCode, fixtureZipCode, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.PostCode: expected masked, got %q", ba.PostCode)
		}
		if !validateStringMask(ba.Country, fixtureCountry, testConfig.MaxPiiStringLength) {
			t.Errorf("BillingAddress.Country: expected masked, got %q", ba.Country)
		}
	})

	t.Run("order", func(t *testing.T) {
		o := result.Orders[0]
		if !validateStringMask(o.OrderID, fixtureOrderID, testConfig.MaxPiiStringLength) {
			t.Errorf("OrderID: expected masked, got %q", o.OrderID)
		}
		if o.Amount != 0 {
			t.Errorf("Amount: expected masked to 0, got %f", o.Amount)
		}
		if !validateStringMask(o.Currency, fixtureCurrency, testConfig.MaxPiiStringLength) {
			t.Errorf("Currency: expected masked, got %q", o.Currency)
		}
		if !validateStringMask(o.Notes, fixtureNotes, testConfig.MaxPiiStringLength) {
			t.Errorf("Notes: expected masked, got %q", o.Notes)
		}
		item := o.Items[0]
		if !validateStringMask(item.ProductID, fixtureProductID, testConfig.MaxPiiStringLength) {
			t.Errorf("ProductID: expected masked, got %q", item.ProductID)
		}
		if !validateStringMask(item.Name, fixtureItemName, testConfig.MaxPiiStringLength) {
			t.Errorf("Name: expected masked, got %q", item.Name)
		}
		if item.Quantity != 0 {
			t.Errorf("Quantity: expected masked to 0, got %d", item.Quantity)
		}
		if item.UnitPrice != 0 {
			t.Errorf("UnitPrice: expected masked to 0, got %f", item.UnitPrice)
		}
	})

	t.Run("device", func(t *testing.T) {
		d := result.Devices["mobile"]
		if !validateStringMask(d.DeviceID, fixtureDeviceID, testConfig.MaxPiiStringLength) {
			t.Errorf("DeviceID: expected masked, got %q", d.DeviceID)
		}
		if !validateStringMask(d.UserAgent, fixtureUserAgent, testConfig.MaxPiiStringLength) {
			t.Errorf("UserAgent: expected masked, got %q", d.UserAgent)
		}
		if !validateStringMask(d.IPAddress, fixtureIPAddress, testConfig.MaxPiiStringLength) {
			t.Errorf("IPAddress: expected masked, got %q", d.IPAddress)
		}
		if !validateStringMask(d.ScreenSize, fixtureScreenSize, testConfig.MaxPiiStringLength) {
			t.Errorf("ScreenSize: expected masked, got %q", d.ScreenSize)
		}
	})
}

func TestMask_PtrChain(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newPtrMaskFixture()
	result := masker.Mask(fixture).(PtrMask)

	t.Run("mask_level", func(t *testing.T) {
		if !validateStringMask(*result.Str, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("Str: expected masked, got %q", *result.Str)
		}
		if *result.Int != 0 {
			t.Errorf("Int: want 0, got %d", *result.Int)
		}
		if *result.Uint != 0 {
			t.Errorf("Uint: want 0, got %d", *result.Uint)
		}
		if *result.Float != 0 {
			t.Errorf("Float: want 0, got %f", *result.Float)
		}
		if *result.Bool != false {
			t.Errorf("Bool: want false, got %v", *result.Bool)
		}
		for i, s := range result.Slice {
			if !validateStringMask(s, fixturePtrStr, testConfig.MaxPiiStringLength) {
				t.Errorf("Slice[%d]: expected masked, got %q", i, s)
			}
		}
		if v := result.MapVal["key"].Value; !validateStringMask(v, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("MapVal[key].Value: expected masked, got %q", v)
		}
		if v := (*result.PMapVal)["key"].Value; !validateStringMask(v, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("PMapVal[key].Value: expected masked, got %q", v)
		}
		if v := result.MapPtr["key"].Value; !validateStringMask(v, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("MapPtr[key].Value: expected masked, got %q", v)
		}
	})

	t.Run("show_level", func(t *testing.T) {
		show := result.Next
		if *show.Str != fixturePtrStr {
			t.Errorf("Str: want %q, got %q", fixturePtrStr, *show.Str)
		}
		if *show.Int != fixturePtrInt {
			t.Errorf("Int: want %d, got %d", fixturePtrInt, *show.Int)
		}
		if *show.Uint != fixturePtrUint {
			t.Errorf("Uint: want %d, got %d", fixturePtrUint, *show.Uint)
		}
		if *show.Float != fixturePtrFloat {
			t.Errorf("Float: want %f, got %f", fixturePtrFloat, *show.Float)
		}
		if *show.Bool != fixturePtrBool {
			t.Errorf("Bool: want %v, got %v", fixturePtrBool, *show.Bool)
		}
		for i, s := range show.Slice {
			if s != fixturePtrStr {
				t.Errorf("Slice[%d]: want %q, got %q", i, fixturePtrStr, s)
			}
		}
		if v := show.MapVal["key"].Value; v != fixturePtrStr {
			t.Errorf("MapVal[key].Value: want %q, got %q", fixturePtrStr, v)
		}
		if v := (*show.PMapVal)["key"].Value; v != fixturePtrStr {
			t.Errorf("PMapVal[key].Value: want %q, got %q", fixturePtrStr, v)
		}
		if v := show.MapPtr["key"].Value; v != fixturePtrStr {
			t.Errorf("MapPtr[key].Value: want %q, got %q", fixturePtrStr, v)
		}
	})

	t.Run("anonymize_level", func(t *testing.T) {
		anon := result.Next.Next
		if !validateAnonymization(*anon.Str, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("Str: want different value of len %d, got %q", len(fixturePtrStr), *anon.Str)
		}
		if *anon.Int == fixturePtrInt {
			t.Errorf("Int: want anonymized value, got original %d", fixturePtrInt)
		}
		if *anon.Uint == fixturePtrUint {
			t.Errorf("Uint: want anonymized value, got original %d", fixturePtrUint)
		}
		if *anon.Float == fixturePtrFloat {
			t.Errorf("Float: want anonymized value, got original %f", fixturePtrFloat)
		}
		for i, s := range anon.Slice {
			if !validateAnonymization(s, fixturePtrStr, testConfig.MaxPiiStringLength) {
				t.Errorf("Slice[%d]: want anonymized string of len %d, got %q", i, len(fixturePtrStr), s)
			}
		}
		if v := anon.MapVal["key"].Value; v == fixturePtrStr {
			t.Errorf("MapVal[key].Value: want anonymized, got original %q", v)
		}
		if v := (*anon.PMapVal)["key"].Value; v == fixturePtrStr {
			t.Errorf("PMapVal[key].Value: want anonymized, got original %q", v)
		}
		if v := anon.MapPtr["key"].Value; v == fixturePtrStr {
			t.Errorf("MapPtr[key].Value: want anonymized, got original %q", v)
		}
	})

	t.Run("notag_level", func(t *testing.T) {
		notag := result.Next.Next.Next
		if !validateStringMask(*notag.Str, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("Str: expected masked, got %q", *notag.Str)
		}
		if *notag.Int != 0 {
			t.Errorf("Int: want 0, got %d", *notag.Int)
		}
		if *notag.Uint != 0 {
			t.Errorf("Uint: want 0, got %d", *notag.Uint)
		}
		if *notag.Float != 0 {
			t.Errorf("Float: want 0, got %f", *notag.Float)
		}
		if *notag.Bool != false {
			t.Errorf("Bool: want false, got %v", *notag.Bool)
		}
		for i, s := range notag.Slice {
			if !validateStringMask(s, fixturePtrStr, testConfig.MaxPiiStringLength) {
				t.Errorf("Slice[%d]: expected masked, got %q", i, s)
			}
		}
		if v := notag.MapVal["key"].Value; !validateStringMask(v, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("MapVal[key].Value: expected masked, got %q", v)
		}
		if v := (*notag.PMapVal)["key"].Value; !validateStringMask(v, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("PMapVal[key].Value: expected masked, got %q", v)
		}
		if v := notag.MapPtr["key"].Value; !validateStringMask(v, fixturePtrStr, testConfig.MaxPiiStringLength) {
			t.Errorf("MapPtr[key].Value: expected masked, got %q", v)
		}
	})
}

func TestMask_UnexportedFields(t *testing.T) {
	masker := newTestMasker(t)

	t.Run("mask", func(t *testing.T) {
		input := newUnexportedFieldsFixture()
		result := masker.Mask(input).(UnexportedFields)
		if result.name != "" || result.age != 0 || result.balance != 0 {
			t.Errorf("expected all fields to remain zero, got name=%q age=%d balance=%f", result.name, result.age, result.balance)
		}
	})

	t.Run("show", func(t *testing.T) {
		input := newUnexportedFieldsShowFixture()
		result := masker.Mask(input).(UnexportedFieldsShow)
		if result.name != "" || result.age != 0 || result.balance != 0 {
			t.Errorf("expected all fields to remain zero, got name=%q age=%d balance=%f", result.name, result.age, result.balance)
		}
	})

	t.Run("anonymize", func(t *testing.T) {
		input := newUnexportedFieldsAnonymizeFixture()
		result := masker.Mask(input).(UnexportedFieldsAnonymize)
		if result.name != "" || result.age != 0 || result.balance != 0 {
			t.Errorf("expected all fields to remain zero, got name=%q age=%d balance=%f", result.name, result.age, result.balance)
		}
	})

	t.Run("no_tag", func(t *testing.T) {
		input := newUnexportedFieldsNoTagFixture()
		result := masker.Mask(input).(UnexportedFieldsNoTag)
		if result.name != "" || result.age != 0 || result.balance != 0 {
			t.Errorf("expected all fields to remain zero, got name=%q age=%d balance=%f", result.name, result.age, result.balance)
		}
	})
}

func BenchmarkMask(b *testing.B) {
	masker := newTestMasker(b)

	b.Run("mask", func(b *testing.B) {
		fixture := newPersonFixture()
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(fixture)
		}
	})

	b.Run("show", func(b *testing.B) {
		fixture := newPersonShowFixture()
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(fixture)
		}
	})

	b.Run("anonymize", func(b *testing.B) {
		fixture := newPersonAnonymizeFixture()
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(fixture)
		}
	})

	b.Run("no_tag", func(b *testing.B) {
		fixture := newPersonNoTagFixture()
		b.ReportAllocs()
		b.ResetTimer()
		for range b.N {
			masker.Mask(fixture)
		}
	})
}
