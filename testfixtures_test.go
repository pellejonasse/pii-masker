package piimasker_test

import "github.com/brianvoe/gofakeit/v6"

var (
	fixtureFirstName  = gofakeit.FirstName()
	fixtureLastName   = gofakeit.LastName()
	fixtureAge        = gofakeit.IntRange(18, 80)
	fixtureEmail      = gofakeit.Email()
	fixturePhone      = gofakeit.Phone()
	fixtureAltEmail   = gofakeit.Email()
	fixtureStreet     = gofakeit.Street()
	fixtureCity       = gofakeit.City()
	fixtureCountry    = gofakeit.Country()
	fixtureZipCode    = gofakeit.Zip()
	fixtureLatitude   = float64(gofakeit.Latitude())
	fixtureLongitude  = float64(gofakeit.Longitude())
	fixtureAltitude   = gofakeit.Float64Range(0, 1000)
	fixtureAccuracy   = gofakeit.Float64Range(1, 100)
	fixtureCardNumber = gofakeit.CreditCardNumber(nil)
	fixtureCVV        = gofakeit.Numerify("###")
	fixtureExpiry     = gofakeit.CreditCardExp()
	fixtureHolderName = gofakeit.Name()
	fixtureOrderID    = gofakeit.UUID()
	fixtureAmount     = gofakeit.Float64Range(1, 9999)
	fixtureCurrency   = gofakeit.CurrencyShort()
	fixtureNotes      = gofakeit.Sentence(5)
	fixtureProductID  = gofakeit.UUID()
	fixtureItemName   = gofakeit.ProductName()
	fixtureQuantity   = gofakeit.IntRange(1, 10)
	fixtureUnitPrice  = gofakeit.Float64Range(1, 500)
	fixtureDeviceID   = gofakeit.UUID()
	fixtureUserAgent  = gofakeit.UserAgent()
	fixtureIPAddress  = gofakeit.IPv4Address()
	fixtureScreenSize = gofakeit.LoremIpsumSentence(2)
)

// -- Raw constructors --

func newPersonFixture() Person {
	return Person{
		FirstName: fixtureFirstName,
		LastName:  fixtureLastName,
		Age:       fixtureAge,
		IsActive:  true,
		Contact: ContactInfo{
			Email:    fixtureEmail,
			Phone:    fixturePhone,
			AltEmail: fixtureAltEmail,
			Address: Address{
				Street:  fixtureStreet,
				City:    fixtureCity,
				Country: fixtureCountry,
				ZipCode: fixtureZipCode,
				Coordinates: GeoCoordinates{
					Latitude:  fixtureLatitude,
					Longitude: fixtureLongitude,
					Altitude:  fixtureAltitude,
					Accuracy:  fixtureAccuracy,
				},
			},
		},
		PaymentMethods: []PaymentMethod{
			{
				CardNumber: fixtureCardNumber,
				CVV:        fixtureCVV,
				Expiry:     fixtureExpiry,
				HolderName: fixtureHolderName,
				IsDefault:  true,
				BillingAddress: CardBillingAddress{
					Street:   fixtureStreet,
					City:     fixtureCity,
					PostCode: fixtureZipCode,
					Country:  fixtureCountry,
				},
			},
		},
		Orders: []Order{
			{
				OrderID:  fixtureOrderID,
				Amount:   fixtureAmount,
				Currency: fixtureCurrency,
				Notes:    fixtureNotes,
				Items: []OrderItem{
					{
						ProductID: fixtureProductID,
						Name:      fixtureItemName,
						Quantity:  fixtureQuantity,
						UnitPrice: fixtureUnitPrice,
					},
				},
			},
		},
		Devices: map[string]DeviceInfo{
			"mobile": {
				DeviceID:   fixtureDeviceID,
				UserAgent:  fixtureUserAgent,
				IPAddress:  fixtureIPAddress,
				ScreenSize: fixtureScreenSize,
			},
		},
	}
}

func newPersonShowFixture() PersonShow {
	return PersonShow{
		FirstName: fixtureFirstName,
		LastName:  fixtureLastName,
		Age:       fixtureAge,
		IsActive:  true,
		Contact: ContactInfoShow{
			Email:    fixtureEmail,
			Phone:    fixturePhone,
			AltEmail: fixtureAltEmail,
			Address: AddressShow{
				Street:  fixtureStreet,
				City:    fixtureCity,
				Country: fixtureCountry,
				ZipCode: fixtureZipCode,
				Coordinates: GeoCoordinatesShow{
					Latitude:  fixtureLatitude,
					Longitude: fixtureLongitude,
					Altitude:  fixtureAltitude,
					Accuracy:  fixtureAccuracy,
				},
			},
		},
		PaymentMethods: []PaymentMethodShow{
			{
				CardNumber: fixtureCardNumber,
				CVV:        fixtureCVV,
				Expiry:     fixtureExpiry,
				HolderName: fixtureHolderName,
				IsDefault:  true,
				BillingAddress: CardBillingAddressShow{
					Street:   fixtureStreet,
					City:     fixtureCity,
					PostCode: fixtureZipCode,
					Country:  fixtureCountry,
				},
			},
		},
		Orders: []OrderShow{
			{
				OrderID:  fixtureOrderID,
				Amount:   fixtureAmount,
				Currency: fixtureCurrency,
				Notes:    fixtureNotes,
				Items: []OrderItemShow{
					{
						ProductID: fixtureProductID,
						Name:      fixtureItemName,
						Quantity:  fixtureQuantity,
						UnitPrice: fixtureUnitPrice,
					},
				},
			},
		},
		Devices: map[string]DeviceInfoShow{
			"mobile": {
				DeviceID:   fixtureDeviceID,
				UserAgent:  fixtureUserAgent,
				IPAddress:  fixtureIPAddress,
				ScreenSize: fixtureScreenSize,
			},
		},
	}
}

func newPersonAnonymizeFixture() PersonAnonymize {
	return PersonAnonymize{
		FirstName: fixtureFirstName,
		LastName:  fixtureLastName,
		Age:       fixtureAge,
		IsActive:  true,
		Contact: ContactInfoAnonymize{
			Email:    fixtureEmail,
			Phone:    fixturePhone,
			AltEmail: fixtureAltEmail,
			Address: AddressAnonymize{
				Street:  fixtureStreet,
				City:    fixtureCity,
				Country: fixtureCountry,
				ZipCode: fixtureZipCode,
				Coordinates: GeoCoordinatesAnonymize{
					Latitude:  fixtureLatitude,
					Longitude: fixtureLongitude,
					Altitude:  fixtureAltitude,
					Accuracy:  fixtureAccuracy,
				},
			},
		},
		PaymentMethods: []PaymentMethodAnonymize{
			{
				CardNumber: fixtureCardNumber,
				CVV:        fixtureCVV,
				Expiry:     fixtureExpiry,
				HolderName: fixtureHolderName,
				IsDefault:  true,
				BillingAddress: CardBillingAddressAnonymize{
					Street:   fixtureStreet,
					City:     fixtureCity,
					PostCode: fixtureZipCode,
					Country:  fixtureCountry,
				},
			},
		},
		Orders: []OrderAnonymize{
			{
				OrderID:  fixtureOrderID,
				Amount:   fixtureAmount,
				Currency: fixtureCurrency,
				Notes:    fixtureNotes,
				Items: []OrderItemAnonymize{
					{
						ProductID: fixtureProductID,
						Name:      fixtureItemName,
						Quantity:  fixtureQuantity,
						UnitPrice: fixtureUnitPrice,
					},
				},
			},
		},
		Devices: map[string]DeviceInfoAnonymize{
			"mobile": {
				DeviceID:   fixtureDeviceID,
				UserAgent:  fixtureUserAgent,
				IPAddress:  fixtureIPAddress,
				ScreenSize: fixtureScreenSize,
			},
		},
	}
}

func newPersonNoTagFixture() PersonNoTag {
	return PersonNoTag{
		FirstName: fixtureFirstName,
		LastName:  fixtureLastName,
		Age:       fixtureAge,
		IsActive:  true,
		Contact: ContactInfoNoTag{
			Email:    fixtureEmail,
			Phone:    fixturePhone,
			AltEmail: fixtureAltEmail,
			Address: AddressNoTag{
				Street:  fixtureStreet,
				City:    fixtureCity,
				Country: fixtureCountry,
				ZipCode: fixtureZipCode,
				Coordinates: GeoCoordinatesNoTag{
					Latitude:  fixtureLatitude,
					Longitude: fixtureLongitude,
					Altitude:  fixtureAltitude,
					Accuracy:  fixtureAccuracy,
				},
			},
		},
		PaymentMethods: []PaymentMethodNoTag{
			{
				CardNumber: fixtureCardNumber,
				CVV:        fixtureCVV,
				Expiry:     fixtureExpiry,
				HolderName: fixtureHolderName,
				IsDefault:  true,
				BillingAddress: CardBillingAddressNoTag{
					Street:   fixtureStreet,
					City:     fixtureCity,
					PostCode: fixtureZipCode,
					Country:  fixtureCountry,
				},
			},
		},
		Orders: []OrderNoTag{
			{
				OrderID:  fixtureOrderID,
				Amount:   fixtureAmount,
				Currency: fixtureCurrency,
				Notes:    fixtureNotes,
				Items: []OrderItemNoTag{
					{
						ProductID: fixtureProductID,
						Name:      fixtureItemName,
						Quantity:  fixtureQuantity,
						UnitPrice: fixtureUnitPrice,
					},
				},
			},
		},
		Devices: map[string]DeviceInfoNoTag{
			"mobile": {
				DeviceID:   fixtureDeviceID,
				UserAgent:  fixtureUserAgent,
				IPAddress:  fixtureIPAddress,
				ScreenSize: fixtureScreenSize,
			},
		},
	}
}

// Pointer fixture data
var (
	fixturePtrStr   = gofakeit.SentenceSimple()
	fixturePtrInt   = gofakeit.IntRange(1, 10000)
	fixturePtrUint  = uint(gofakeit.UintRange(1, 10000))
	fixturePtrFloat = gofakeit.Float64Range(1, 10000)
	fixturePtrBool  = gofakeit.Bool()
)

func ptrStr(s string) *string     { return &s }
func ptrInt(i int) *int           { return &i }
func ptrUint(u uint) *uint        { return &u }
func ptrFloat(f float64) *float64 { return &f }
func ptrBool(b bool) *bool        { return &b }

func newPtrMaskFixture() PtrMask {
	fields := func() (map[string]PtrMapValue, *map[string]PtrMapValue, map[string]*PtrMapValue) {
		m := map[string]PtrMapValue{"key": {Value: fixturePtrStr}}
		mp := map[string]*PtrMapValue{"key": {Value: fixturePtrStr}}
		return m, &m, mp
	}

	noTagM, noTagPM, noTagMP := fields()
	anonM, anonPM, anonMP := fields()
	showM, showPM, showMP := fields()
	maskM, maskPM, maskMP := fields()

	return PtrMask{
		Str:     ptrStr(fixturePtrStr),
		Int:     ptrInt(fixturePtrInt),
		Uint:    ptrUint(fixturePtrUint),
		Float:   ptrFloat(fixturePtrFloat),
		Bool:    ptrBool(fixturePtrBool),
		Slice:   []string{fixturePtrStr, fixturePtrStr},
		MapVal:  maskM,
		PMapVal: maskPM,
		MapPtr:  maskMP,
		Next: &PtrShow{
			Str:     ptrStr(fixturePtrStr),
			Int:     ptrInt(fixturePtrInt),
			Uint:    ptrUint(fixturePtrUint),
			Float:   ptrFloat(fixturePtrFloat),
			Bool:    ptrBool(fixturePtrBool),
			Slice:   []string{fixturePtrStr, fixturePtrStr},
			MapVal:  showM,
			PMapVal: showPM,
			MapPtr:  showMP,
			Next: &PtrAnonymize{
				Str:     ptrStr(fixturePtrStr),
				Int:     ptrInt(fixturePtrInt),
				Uint:    ptrUint(fixturePtrUint),
				Float:   ptrFloat(fixturePtrFloat),
				Bool:    ptrBool(fixturePtrBool),
				Slice:   []string{fixturePtrStr, fixturePtrStr},
				MapVal:  anonM,
				PMapVal: anonPM,
				MapPtr:  anonMP,
				Next: &PtrNoTag{
					Str:     ptrStr(fixturePtrStr),
					Int:     ptrInt(fixturePtrInt),
					Uint:    ptrUint(fixturePtrUint),
					Float:   ptrFloat(fixturePtrFloat),
					Bool:    ptrBool(fixturePtrBool),
					Slice:   []string{fixturePtrStr, fixturePtrStr},
					MapVal:  noTagM,
					PMapVal: noTagPM,
					MapPtr:  noTagMP,
				},
			},
		},
	}
}

// Unexported field fixture data
var (
	fixtureUnexportedName    = gofakeit.Name()
	fixtureUnexportedAge     = gofakeit.IntRange(18, 80)
	fixtureUnexportedBalance = gofakeit.Float64Range(1, 9999)
)

func newUnexportedFieldsFixture() UnexportedFields {
	return UnexportedFields{name: fixtureUnexportedName, age: fixtureUnexportedAge, balance: fixtureUnexportedBalance}
}

func newUnexportedFieldsShowFixture() UnexportedFieldsShow {
	return UnexportedFieldsShow{name: fixtureUnexportedName, age: fixtureUnexportedAge, balance: fixtureUnexportedBalance}
}

func newUnexportedFieldsAnonymizeFixture() UnexportedFieldsAnonymize {
	return UnexportedFieldsAnonymize{name: fixtureUnexportedName, age: fixtureUnexportedAge, balance: fixtureUnexportedBalance}
}

func newUnexportedFieldsNoTagFixture() UnexportedFieldsNoTag {
	return UnexportedFieldsNoTag{name: fixtureUnexportedName, age: fixtureUnexportedAge, balance: fixtureUnexportedBalance}
}
