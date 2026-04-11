package piimasker_test

// raw fixture data — shared values used across all variants, and their constructors
const (
	fixtureFirstName  = "John"
	fixtureLastName   = "Smith"
	fixtureAge        = 30
	fixtureEmail      = "john.smith@example.com"
	fixturePhone      = "+44 7700 900000"
	fixtureAltEmail   = "j.smith@work.com"
	fixtureStreet     = "10 Downing St"
	fixtureCity       = "London"
	fixtureCountry    = "UK"
	fixtureZipCode    = "SW1A 2AA"
	fixtureLatitude   = 51.5074
	fixtureLongitude  = -0.1278
	fixtureAltitude   = 11.0
	fixtureAccuracy   = 5.0
	fixtureCardNumber = "4111111111111111"
	fixtureCVV        = "123"
	fixtureExpiry     = "12/28"
	fixtureHolderName = "John Smith"
	fixtureOrderID    = "ORD-001"
	fixtureAmount     = 99.99
	fixtureCurrency   = "GBP"
	fixtureNotes      = "Leave at door"
	fixtureProductID  = "PROD-001"
	fixtureItemName   = "Widget"
	fixtureQuantity   = 2
	fixtureUnitPrice  = 49.99
	fixtureDeviceID   = "DEV-001"
	fixtureUserAgent  = "Mozilla/5.0"
	fixtureIPAddress  = "192.168.1.1"
	fixtureScreenSize = "1920x1080"
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
