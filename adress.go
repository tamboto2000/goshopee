package goshopee

type UserAddresses struct {
	Addresses             []Address     `json:"addresses,omitempty"`
	PickupAddressID       int           `json:"pickup_address_id,omitempty"`
	DeliveryAddressID     int           `json:"delivery_address_id,omitempty"`
	ReturnAddressID       int           `json:"return_address_id,omitempty"`
	SellerReturnAddressID int           `json:"seller_return_address_id,omitempty"`
	LogisticAddresses     []interface{} `json:"logistic_addresses,omitempty"`
}

type Address struct {
	LogisticsStatus    int             `json:"logistics_status,omitempty"`
	Mtime              int             `json:"mtime,omitempty"`
	Icno               string          `json:"icno,omitempty"`
	ID                 int             `json:"id,omitempty"`
	City               string          `json:"city,omitempty"`
	District           string          `json:"district,omitempty"`
	Zipcode            string          `json:"zipcode,omitempty"`
	Label              string          `json:"label,omitempty"`
	State              string          `json:"state,omitempty"`
	GeoString          string          `json:"geoString,omitempty"`
	Status             int             `json:"status,omitempty"`
	Deftime            int             `json:"deftime,omitempty"`
	FullAddress        string          `json:"full_address,omitempty"`
	Phone              string          `json:"phone,omitempty"`
	FormattedAddress   string          `json:"formattedAddress,omitempty"`
	AddressInstruction string          `json:"address_instruction,omitempty"`
	Address            string          `json:"address,omitempty"`
	Extinfo            *AddressExtinfo `json:"extinfo,omitempty"`
	Name               string          `json:"name,omitempty"`
	Town               string          `json:"town,omitempty"`
	Ctime              int             `json:"ctime,omitempty"`
	Country            string          `json:"country,omitempty"`
	Userid             int             `json:"userid,omitempty"`
	Geoinfo            *Geoinfo        `json:"geoinfo,omitempty"`
}

type AddressExtinfo struct {
	Geoinfo                 string `json:"geoinfo,omitempty"`
	PreferredDeliveryOption int    `json:"preferred_delivery_option,omitempty"`
	AddressInstruction      string `json:"address_instruction,omitempty"`
	Label                   string `json:"label,omitempty"`
}

type Geoinfo struct {
	Region           *Region `json:"region,omitempty"`
	FormattedAddress string  `json:"formattedAddress,omitempty"`
	Country          string  `json:"country,omitempty"`
	Admin1           string  `json:"admin1,omitempty"`
	Admin3           string  `json:"admin3,omitempty"`
	Admin2           string  `json:"admin2,omitempty"`
	Admin4           string  `json:"admin4,omitempty"`
	StreetNumber     string  `json:"streetNumber,omitempty"`
	PostalCode       string  `json:"postalCode,omitempty"`
	StreetName       string  `json:"streetName,omitempty"`
	PlaceID          string  `json:"placeId,omitempty"`
}

type Region struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
