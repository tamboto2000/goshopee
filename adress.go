package goshopee

type UserAddresses struct {
	Addresses             []Address     `json:"addresses"`
	PickupAddressID       int64         `json:"pickup_address_id"`
	DeliveryAddressID     int64         `json:"delivery_address_id"`
	ReturnAddressID       int64         `json:"return_address_id"`
	SellerReturnAddressID int64         `json:"seller_return_address_id"`
	LogisticAddresses     []interface{} `json:"logistic_addresses"`
}

type Address struct {
	LogisticsStatus    int64          `json:"logistics_status"`
	Mtime              int64          `json:"mtime"`
	Icno               string         `json:"icno"`
	ID                 int64          `json:"id"`
	City               string         `json:"city"`
	District           string         `json:"district"`
	Zipcode            string         `json:"zipcode"`
	Label              string         `json:"label"`
	State              string         `json:"state"`
	GeoString          string         `json:"geoString"`
	Status             int64          `json:"status"`
	Deftime            int64          `json:"deftime"`
	FullAddress        string         `json:"full_address"`
	Phone              string         `json:"phone"`
	FormattedAddress   string         `json:"formattedAddress"`
	AddressInstruction string         `json:"address_instruction"`
	Address            string         `json:"address"`
	Extinfo            AddressExtinfo `json:"extinfo"`
	Name               string         `json:"name"`
	Town               string         `json:"town"`
	Ctime              int64          `json:"ctime"`
	Country            string         `json:"country"`
	Userid             int64          `json:"userid"`
	Geoinfo            Geoinfo        `json:"geoinfo"`
}

type AddressExtinfo struct {
	Geoinfo                 string  `json:"geoinfo"`
	PreferredDeliveryOption *int64  `json:"preferred_delivery_option,omitempty"`
	AddressInstruction      *string `json:"address_instruction,omitempty"`
	Label                   *string `json:"label,omitempty"`
}

type Geoinfo struct {
	Region           Region  `json:"region"`
	FormattedAddress string  `json:"formattedAddress"`
	Country          *string `json:"country,omitempty"`
	Admin1           *string `json:"admin1,omitempty"`
	Admin3           *string `json:"admin3,omitempty"`
	Admin2           *string `json:"admin2,omitempty"`
	Admin4           *string `json:"admin4,omitempty"`
	StreetNumber     *string `json:"streetNumber,omitempty"`
	PostalCode       *string `json:"postalCode,omitempty"`
	StreetName       *string `json:"streetName,omitempty"`
	PlaceID          *string `json:"placeId,omitempty"`
}

type Region struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
