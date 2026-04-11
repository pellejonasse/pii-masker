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

func TestMask_Integration_MaskTag(t *testing.T) {
	m := newTestMasker(t)
	result := m.Mask(newPersonFixture()).(Person)

	t.Run("top_level", func(t *testing.T) {
		if result.FirstName != "" {
			t.Errorf("FirstName: expected masked, got %q", result.FirstName)
		}
		if result.LastName != "" {
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
		if c.Email != "" {
			t.Errorf("Email: expected masked, got %q", c.Email)
		}
		if c.Phone != "" {
			t.Errorf("Phone: expected masked, got %q", c.Phone)
		}
		if c.AltEmail != "" {
			t.Errorf("AltEmail: expected masked, got %q", c.AltEmail)
		}
	})

	t.Run("address", func(t *testing.T) {
		a := result.Contact.Address
		if a.Street != "" {
			t.Errorf("Street: expected masked, got %q", a.Street)
		}
		if a.City != "" {
			t.Errorf("City: expected masked, got %q", a.City)
		}
		if a.Country != "" {
			t.Errorf("Country: expected masked, got %q", a.Country)
		}
		if a.ZipCode != "" {
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
		if pm.CardNumber != "" {
			t.Errorf("CardNumber: expected masked, got %q", pm.CardNumber)
		}
		if pm.CVV != "" {
			t.Errorf("CVV: expected masked, got %q", pm.CVV)
		}
		if pm.Expiry != "" {
			t.Errorf("Expiry: expected masked, got %q", pm.Expiry)
		}
		if pm.HolderName != "" {
			t.Errorf("HolderName: expected masked, got %q", pm.HolderName)
		}
		if pm.IsDefault != false {
			t.Errorf("IsDefault: expected masked to false, got %v", pm.IsDefault)
		}
		ba := pm.BillingAddress
		if ba.Street != "" {
			t.Errorf("BillingAddress.Street: expected masked, got %q", ba.Street)
		}
		if ba.City != "" {
			t.Errorf("BillingAddress.City: expected masked, got %q", ba.City)
		}
		if ba.PostCode != "" {
			t.Errorf("BillingAddress.PostCode: expected masked, got %q", ba.PostCode)
		}
		if ba.Country != "" {
			t.Errorf("BillingAddress.Country: expected masked, got %q", ba.Country)
		}
	})

	t.Run("order", func(t *testing.T) {
		o := result.Orders[0]
		if o.OrderID != "" {
			t.Errorf("OrderID: expected masked, got %q", o.OrderID)
		}
		if o.Amount != 0 {
			t.Errorf("Amount: expected masked to 0, got %f", o.Amount)
		}
		if o.Currency != "" {
			t.Errorf("Currency: expected masked, got %q", o.Currency)
		}
		if o.Notes != "" {
			t.Errorf("Notes: expected masked, got %q", o.Notes)
		}
		item := o.Items[0]
		if item.ProductID != "" {
			t.Errorf("ProductID: expected masked, got %q", item.ProductID)
		}
		if item.Name != "" {
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
		if d.DeviceID != "" {
			t.Errorf("DeviceID: expected masked, got %q", d.DeviceID)
		}
		if d.UserAgent != "" {
			t.Errorf("UserAgent: expected masked, got %q", d.UserAgent)
		}
		if d.IPAddress != "" {
			t.Errorf("IPAddress: expected masked, got %q", d.IPAddress)
		}
		if d.ScreenSize != "" {
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
		if result.FirstName == fixture.FirstName || len(result.FirstName) != len(fixture.FirstName) {
			t.Errorf("FirstName: expected anonymized same-length string, got %q", result.FirstName)
		}
		if result.LastName == fixture.LastName || len(result.LastName) != len(fixture.LastName) {
			t.Errorf("LastName: expected anonymized same-length string, got %q", result.LastName)
		}
		if result.Age == 0 {
			t.Errorf("Age: expected anonymized non-zero, got 0")
		}
		// IsActive is a bool — skipped; non-deterministic single-run assertion, covered by TestMask_AnonymizeTag/bool
	})

	t.Run("contact", func(t *testing.T) {
		c := result.Contact
		if c.Email == fixture.Contact.Email || len(c.Email) != len(fixture.Contact.Email) {
			t.Errorf("Email: expected anonymized same-length string, got %q", c.Email)
		}
		if c.Phone == fixture.Contact.Phone || len(c.Phone) != len(fixture.Contact.Phone) {
			t.Errorf("Phone: expected anonymized same-length string, got %q", c.Phone)
		}
		if c.AltEmail == fixture.Contact.AltEmail || len(c.AltEmail) != len(fixture.Contact.AltEmail) {
			t.Errorf("AltEmail: expected anonymized same-length string, got %q", c.AltEmail)
		}
	})

	t.Run("address", func(t *testing.T) {
		a := result.Contact.Address
		fa := fixture.Contact.Address
		if a.Street == fa.Street || len(a.Street) != len(fa.Street) {
			t.Errorf("Street: expected anonymized same-length string, got %q", a.Street)
		}
		if a.City == fa.City || len(a.City) != len(fa.City) {
			t.Errorf("City: expected anonymized same-length string, got %q", a.City)
		}
		if a.Country == fa.Country || len(a.Country) != len(fa.Country) {
			t.Errorf("Country: expected anonymized same-length string, got %q", a.Country)
		}
		if a.ZipCode == fa.ZipCode || len(a.ZipCode) != len(fa.ZipCode) {
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
		if pm.CardNumber == fpm.CardNumber || len(pm.CardNumber) != len(fpm.CardNumber) {
			t.Errorf("CardNumber: expected anonymized same-length string, got %q", pm.CardNumber)
		}
		if pm.CVV == fpm.CVV || len(pm.CVV) != len(fpm.CVV) {
			t.Errorf("CVV: expected anonymized same-length string, got %q", pm.CVV)
		}
		if pm.Expiry == fpm.Expiry || len(pm.Expiry) != len(fpm.Expiry) {
			t.Errorf("Expiry: expected anonymized same-length string, got %q", pm.Expiry)
		}
		if pm.HolderName == fpm.HolderName || len(pm.HolderName) != len(fpm.HolderName) {
			t.Errorf("HolderName: expected anonymized same-length string, got %q", pm.HolderName)
		}
		// IsDefault is a bool — skipped; non-deterministic single-run assertion, covered by TestMask_AnonymizeTag/bool
		ba := pm.BillingAddress
		fba := fpm.BillingAddress
		if ba.Street == fba.Street || len(ba.Street) != len(fba.Street) {
			t.Errorf("BillingAddress.Street: expected anonymized same-length string, got %q", ba.Street)
		}
		if ba.City == fba.City || len(ba.City) != len(fba.City) {
			t.Errorf("BillingAddress.City: expected anonymized same-length string, got %q", ba.City)
		}
		if ba.PostCode == fba.PostCode || len(ba.PostCode) != len(fba.PostCode) {
			t.Errorf("BillingAddress.PostCode: expected anonymized same-length string, got %q", ba.PostCode)
		}
		if ba.Country == fba.Country || len(ba.Country) != len(fba.Country) {
			t.Errorf("BillingAddress.Country: expected anonymized same-length string, got %q", ba.Country)
		}
	})

	t.Run("order", func(t *testing.T) {
		o := result.Orders[0]
		fo := fixture.Orders[0]
		if o.OrderID == fo.OrderID || len(o.OrderID) != len(fo.OrderID) {
			t.Errorf("OrderID: expected anonymized same-length string, got %q", o.OrderID)
		}
		if o.Amount == 0 {
			t.Errorf("Amount: expected anonymized non-zero, got 0")
		}
		if o.Currency == fo.Currency || len(o.Currency) != len(fo.Currency) {
			t.Errorf("Currency: expected anonymized same-length string, got %q", o.Currency)
		}
		if o.Notes == fo.Notes || len(o.Notes) != len(fo.Notes) {
			t.Errorf("Notes: expected anonymized same-length string, got %q", o.Notes)
		}
		item := o.Items[0]
		fitem := fo.Items[0]
		if item.ProductID == fitem.ProductID || len(item.ProductID) != len(fitem.ProductID) {
			t.Errorf("ProductID: expected anonymized same-length string, got %q", item.ProductID)
		}
		if item.Name == fitem.Name || len(item.Name) != len(fitem.Name) {
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
		if d.DeviceID == fd.DeviceID || len(d.DeviceID) != len(fd.DeviceID) {
			t.Errorf("DeviceID: expected anonymized same-length string, got %q", d.DeviceID)
		}
		if d.UserAgent == fd.UserAgent || len(d.UserAgent) != len(fd.UserAgent) {
			t.Errorf("UserAgent: expected anonymized same-length string, got %q", d.UserAgent)
		}
		if d.IPAddress == fd.IPAddress || len(d.IPAddress) != len(fd.IPAddress) {
			t.Errorf("IPAddress: expected anonymized same-length string, got %q", d.IPAddress)
		}
		if d.ScreenSize == fd.ScreenSize || len(d.ScreenSize) != len(fd.ScreenSize) {
			t.Errorf("ScreenSize: expected anonymized same-length string, got %q", d.ScreenSize)
		}
	})
}

func TestMask_Integration_NoTag(t *testing.T) {
	m := newTestMasker(t)
	result := m.Mask(newPersonNoTagFixture()).(PersonNoTag)

	t.Run("top_level", func(t *testing.T) {
		if result.FirstName != "" {
			t.Errorf("FirstName: expected masked, got %q", result.FirstName)
		}
		if result.LastName != "" {
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
		if c.Email != "" {
			t.Errorf("Email: expected masked, got %q", c.Email)
		}
		if c.Phone != "" {
			t.Errorf("Phone: expected masked, got %q", c.Phone)
		}
		if c.AltEmail != "" {
			t.Errorf("AltEmail: expected masked, got %q", c.AltEmail)
		}
	})

	t.Run("address", func(t *testing.T) {
		a := result.Contact.Address
		if a.Street != "" {
			t.Errorf("Street: expected masked, got %q", a.Street)
		}
		if a.City != "" {
			t.Errorf("City: expected masked, got %q", a.City)
		}
		if a.Country != "" {
			t.Errorf("Country: expected masked, got %q", a.Country)
		}
		if a.ZipCode != "" {
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
		if pm.CardNumber != "" {
			t.Errorf("CardNumber: expected masked, got %q", pm.CardNumber)
		}
		if pm.CVV != "" {
			t.Errorf("CVV: expected masked, got %q", pm.CVV)
		}
		if pm.Expiry != "" {
			t.Errorf("Expiry: expected masked, got %q", pm.Expiry)
		}
		if pm.HolderName != "" {
			t.Errorf("HolderName: expected masked, got %q", pm.HolderName)
		}
		if pm.IsDefault != false {
			t.Errorf("IsDefault: expected masked to false, got %v", pm.IsDefault)
		}
		ba := pm.BillingAddress
		if ba.Street != "" {
			t.Errorf("BillingAddress.Street: expected masked, got %q", ba.Street)
		}
		if ba.City != "" {
			t.Errorf("BillingAddress.City: expected masked, got %q", ba.City)
		}
		if ba.PostCode != "" {
			t.Errorf("BillingAddress.PostCode: expected masked, got %q", ba.PostCode)
		}
		if ba.Country != "" {
			t.Errorf("BillingAddress.Country: expected masked, got %q", ba.Country)
		}
	})

	t.Run("order", func(t *testing.T) {
		o := result.Orders[0]
		if o.OrderID != "" {
			t.Errorf("OrderID: expected masked, got %q", o.OrderID)
		}
		if o.Amount != 0 {
			t.Errorf("Amount: expected masked to 0, got %f", o.Amount)
		}
		if o.Currency != "" {
			t.Errorf("Currency: expected masked, got %q", o.Currency)
		}
		if o.Notes != "" {
			t.Errorf("Notes: expected masked, got %q", o.Notes)
		}
		item := o.Items[0]
		if item.ProductID != "" {
			t.Errorf("ProductID: expected masked, got %q", item.ProductID)
		}
		if item.Name != "" {
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
		if d.DeviceID != "" {
			t.Errorf("DeviceID: expected masked, got %q", d.DeviceID)
		}
		if d.UserAgent != "" {
			t.Errorf("UserAgent: expected masked, got %q", d.UserAgent)
		}
		if d.IPAddress != "" {
			t.Errorf("IPAddress: expected masked, got %q", d.IPAddress)
		}
		if d.ScreenSize != "" {
			t.Errorf("ScreenSize: expected masked, got %q", d.ScreenSize)
		}
	})
}

func TestMask_PtrChain(t *testing.T) {
	masker := newTestMasker(t)
	fixture := newPtrMaskFixture()
	result := masker.Mask(fixture).(PtrMask)

	t.Run("mask_level", func(t *testing.T) {
		if *result.Str != "" {
			t.Errorf("Str: want %q, got %q", "", *result.Str)
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
			if s != "" {
				t.Errorf("Slice[%d]: want %q, got %q", i, "", s)
			}
		}
		if v := result.MapVal["key"].Value; v != "" {
			t.Errorf("MapVal[key].Value: want %q, got %q", "", v)
		}
		if v := (*result.PMapVal)["key"].Value; v != "" {
			t.Errorf("PMapVal[key].Value: want %q, got %q", "", v)
		}
		if v := result.MapPtr["key"].Value; v != "" {
			t.Errorf("MapPtr[key].Value: want %q, got %q", "", v)
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
		if got := *anon.Str; len(got) != len(fixturePtrStr) || got == fixturePtrStr {
			t.Errorf("Str: want different value of len %d, got %q", len(fixturePtrStr), got)
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
			if len(s) != len(fixturePtrStr) || s == fixturePtrStr {
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
		if *notag.Str != "" {
			t.Errorf("Str: want %q, got %q", "", *notag.Str)
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
			if s != "" {
				t.Errorf("Slice[%d]: want %q, got %q", i, "", s)
			}
		}
		if v := notag.MapVal["key"].Value; v != "" {
			t.Errorf("MapVal[key].Value: want %q, got %q", "", v)
		}
		if v := (*notag.PMapVal)["key"].Value; v != "" {
			t.Errorf("PMapVal[key].Value: want %q, got %q", "", v)
		}
		if v := notag.MapPtr["key"].Value; v != "" {
			t.Errorf("MapPtr[key].Value: want %q, got %q", "", v)
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
