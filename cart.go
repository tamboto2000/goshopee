package goshopee

// Cart contains info about your cart
type Cart struct {
	Data         Data   `json:"data,omitempty"`
	Error        int    `json:"error,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`

	sh *Shopee
}

// Cart get your cart info
func (sh *Shopee) Cart() (*Cart, error) {
	// raw, err := sh.post("")

	return nil, nil
}
