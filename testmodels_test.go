package piimasker_test

// -- Level 3 (deepest) --

type GeoCoordinates struct {
	Latitude  float64 `Pii:"mask"`
	Longitude float64 `Pii:"mask"`
	Altitude  float64 `Pii:"mask"`
	Accuracy  float64 `Pii:"mask"`
}

type CardBillingAddress struct {
	Street   string `Pii:"mask"`
	City     string `Pii:"mask"`
	PostCode string `Pii:"mask"`
	Country  string `Pii:"mask"`
}

// -- Level 2 --

type Address struct {
	Street      string `Pii:"mask"`
	City        string `Pii:"mask"`
	Country     string `Pii:"mask"`
	ZipCode     string `Pii:"mask"`
	Coordinates GeoCoordinates
}

type PaymentMethod struct {
	CardNumber     string `Pii:"mask"`
	CVV            string `Pii:"mask"`
	Expiry         string `Pii:"mask"`
	HolderName     string `Pii:"mask"`
	IsDefault      bool   `Pii:"mask"`
	BillingAddress CardBillingAddress
}

type OrderItem struct {
	ProductID string  `Pii:"mask"`
	Name      string  `Pii:"mask"`
	Quantity  int     `Pii:"mask"`
	UnitPrice float64 `Pii:"mask"`
}

type DeviceInfo struct {
	DeviceID   string `Pii:"mask"`
	UserAgent  string `Pii:"mask"`
	IPAddress  string `Pii:"mask"`
	ScreenSize string `Pii:"mask"`
}

// -- Level 1 --

type ContactInfo struct {
	Email    string `Pii:"mask"`
	Phone    string `Pii:"mask"`
	Address  Address
	AltEmail string `Pii:"mask"`
}

type Order struct {
	OrderID  string  `Pii:"mask"`
	Amount   float64 `Pii:"mask"`
	Currency string  `Pii:"mask"`
	Notes    string  `Pii:"mask"`
	Items    []OrderItem
}

// -- Level 0 (top-level) --

type Person struct {
	FirstName      string `Pii:"mask"`
	LastName       string `Pii:"mask"`
	Age            int    `Pii:"mask"`
	Contact        ContactInfo
	PaymentMethods []PaymentMethod
	Orders         []Order
	Devices        map[string]DeviceInfo
	IsActive       bool `Pii:"mask"`
}

// -- Show variants --

type GeoCoordinatesShow struct {
	Latitude  float64 `Pii:"show"`
	Longitude float64 `Pii:"show"`
	Altitude  float64 `Pii:"show"`
	Accuracy  float64 `Pii:"show"`
}

type CardBillingAddressShow struct {
	Street   string `Pii:"show"`
	City     string `Pii:"show"`
	PostCode string `Pii:"show"`
	Country  string `Pii:"show"`
}

type AddressShow struct {
	Street      string `Pii:"show"`
	City        string `Pii:"show"`
	Country     string `Pii:"show"`
	ZipCode     string `Pii:"show"`
	Coordinates GeoCoordinatesShow
}

type PaymentMethodShow struct {
	CardNumber     string `Pii:"show"`
	CVV            string `Pii:"show"`
	Expiry         string `Pii:"show"`
	HolderName     string `Pii:"show"`
	IsDefault      bool   `Pii:"show"`
	BillingAddress CardBillingAddressShow
}

type OrderItemShow struct {
	ProductID string  `Pii:"show"`
	Name      string  `Pii:"show"`
	Quantity  int     `Pii:"show"`
	UnitPrice float64 `Pii:"show"`
}

type DeviceInfoShow struct {
	DeviceID   string `Pii:"show"`
	UserAgent  string `Pii:"show"`
	IPAddress  string `Pii:"show"`
	ScreenSize string `Pii:"show"`
}

type ContactInfoShow struct {
	Email    string `Pii:"show"`
	Phone    string `Pii:"show"`
	Address  AddressShow
	AltEmail string `Pii:"show"`
}

type OrderShow struct {
	OrderID  string  `Pii:"show"`
	Amount   float64 `Pii:"show"`
	Currency string  `Pii:"show"`
	Notes    string  `Pii:"show"`
	Items    []OrderItemShow
}

type PersonShow struct {
	FirstName      string `Pii:"show"`
	LastName       string `Pii:"show"`
	Age            int    `Pii:"show"`
	Contact        ContactInfoShow
	PaymentMethods []PaymentMethodShow
	Orders         []OrderShow
	Devices        map[string]DeviceInfoShow
	IsActive       bool `Pii:"show"`
}

// -- Anonymize variants --

type GeoCoordinatesAnonymize struct {
	Latitude  float64 `Pii:"anonymize"`
	Longitude float64 `Pii:"anonymize"`
	Altitude  float64 `Pii:"anonymize"`
	Accuracy  float64 `Pii:"anonymize"`
}

type CardBillingAddressAnonymize struct {
	Street   string `Pii:"anonymize"`
	City     string `Pii:"anonymize"`
	PostCode string `Pii:"anonymize"`
	Country  string `Pii:"anonymize"`
}

type AddressAnonymize struct {
	Street      string `Pii:"anonymize"`
	City        string `Pii:"anonymize"`
	Country     string `Pii:"anonymize"`
	ZipCode     string `Pii:"anonymize"`
	Coordinates GeoCoordinatesAnonymize
}

type PaymentMethodAnonymize struct {
	CardNumber     string `Pii:"anonymize"`
	CVV            string `Pii:"anonymize"`
	Expiry         string `Pii:"anonymize"`
	HolderName     string `Pii:"anonymize"`
	IsDefault      bool   `Pii:"anonymize"`
	BillingAddress CardBillingAddressAnonymize
}

type OrderItemAnonymize struct {
	ProductID string  `Pii:"anonymize"`
	Name      string  `Pii:"anonymize"`
	Quantity  int     `Pii:"anonymize"`
	UnitPrice float64 `Pii:"anonymize"`
}

type DeviceInfoAnonymize struct {
	DeviceID   string `Pii:"anonymize"`
	UserAgent  string `Pii:"anonymize"`
	IPAddress  string `Pii:"anonymize"`
	ScreenSize string `Pii:"anonymize"`
}

type ContactInfoAnonymize struct {
	Email    string `Pii:"anonymize"`
	Phone    string `Pii:"anonymize"`
	Address  AddressAnonymize
	AltEmail string `Pii:"anonymize"`
}

type OrderAnonymize struct {
	OrderID  string  `Pii:"anonymize"`
	Amount   float64 `Pii:"anonymize"`
	Currency string  `Pii:"anonymize"`
	Notes    string  `Pii:"anonymize"`
	Items    []OrderItemAnonymize
}

type PersonAnonymize struct {
	FirstName      string `Pii:"anonymize"`
	LastName       string `Pii:"anonymize"`
	Age            int    `Pii:"anonymize"`
	Contact        ContactInfoAnonymize
	PaymentMethods []PaymentMethodAnonymize
	Orders         []OrderAnonymize
	Devices        map[string]DeviceInfoAnonymize
	IsActive       bool `Pii:"anonymize"`
}

// -- No-tag variants (default/inherit behaviour) --

type GeoCoordinatesNoTag struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
	Accuracy  float64
}

type CardBillingAddressNoTag struct {
	Street   string
	City     string
	PostCode string
	Country  string
}

type AddressNoTag struct {
	Street      string
	City        string
	Country     string
	ZipCode     string
	Coordinates GeoCoordinatesNoTag
}

type PaymentMethodNoTag struct {
	CardNumber     string
	CVV            string
	Expiry         string
	HolderName     string
	IsDefault      bool
	BillingAddress CardBillingAddressNoTag
}

type OrderItemNoTag struct {
	ProductID string
	Name      string
	Quantity  int
	UnitPrice float64
}

type DeviceInfoNoTag struct {
	DeviceID   string
	UserAgent  string
	IPAddress  string
	ScreenSize string
}

type ContactInfoNoTag struct {
	Email    string
	Phone    string
	Address  AddressNoTag
	AltEmail string
}

type OrderNoTag struct {
	OrderID  string
	Amount   float64
	Currency string
	Notes    string
	Items    []OrderItemNoTag
}

type PersonNoTag struct {
	FirstName      string
	LastName       string
	Age            int
	Contact        ContactInfoNoTag
	PaymentMethods []PaymentMethodNoTag
	Orders         []OrderNoTag
	Devices        map[string]DeviceInfoNoTag
	IsActive       bool
}
