package goshopee

import "encoding/json"

// Cart contains info about your cart
type Cart struct {
	Data         Data   `json:"data,omitempty"`
	Error        int    `json:"error,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`

	sh *Shopee
}

// Cart get your cart info
func (sh *Shopee) Cart() (*Cart, error) {
	raw, err := sh.post("/v4/cart/get", "https://shopee.co.id/cart/", map[string]interface{}{
		"pre_selected_item_list": make([]interface{}, 0),
	})

	if err != nil {
		return nil, err
	}

	cart := new(Cart)

	return cart, json.Unmarshal(raw, cart)
}
