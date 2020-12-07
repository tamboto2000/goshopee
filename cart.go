package goshopee

import "encoding/json"

// Cart actions
const (
	SelectAllItems   = 4
	DeselectAllItems = 5
)

// Cart contains info about your cart
type Cart struct {
	Data                 Data          `json:"data,omitempty"`
	Error                int           `json:"error,omitempty"`
	ErrorMessage         string        `json:"error_message,omitempty"`
	MessageLevel         MessageLevel  `json:"message_level,omitempty"`
	WarnMessage          interface{}   `json:"warn_message,omitempty"`
	SelectedShopOrderIDS []ShopOrder   `json:"selected_shop_order_ids,omitempty"`
	PlatformVouchers     []interface{} `json:"platform_vouchers,omitempty"`

	sh *Shopee
}

type MessageLevel struct {
	Toast   bool        `json:"toast,omitempty"`
	Popup   bool        `json:"popup,omitempty"`
	Refresh interface{} `json:"refresh,omitempty"`
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

	if err := json.Unmarshal(raw, cart); err != nil {
		return nil, err
	}

	cart.sh = sh

	return cart, nil
}

// // CheckoutAll checkout all items in cart
// func (cart *Cart) CheckoutAll() (*Cart, error) {
// 	body := make(map[string]interface{})

// 	return nil, nil
// }
