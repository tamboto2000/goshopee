package goshopee

import (
	"encoding/json"
	"errors"

	"github.com/tamboto2000/mapper"
)

// Cart actions
const (
	SelectAllItems   = 4
	DeselectAllItems = 5
)

// Cart contains info about your cart
type Cart struct {
	Data                 *Data         `json:"data,omitempty"`
	Error                int           `json:"error,omitempty"`
	ErrorMessage         string        `json:"error_message,omitempty"`
	MessageLevel         *MessageLevel `json:"message_level,omitempty"`
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

// CheckoutAll checkout all items in cart
func (cart *Cart) CheckoutAll() error {
	body := make(map[string]interface{})
	shopOrders := make([]ShopOrder, 0)

	// prepare all items in cart
	if cart.Data.ShopOrders != nil {
		if len(cart.Data.ShopOrders) > 0 {
			for _, shopOrder := range cart.Data.ShopOrders {
				newShopOrder := ShopOrder{
					Shopid:       shopOrder.Shop.Shopid,
					ShopVouchers: make([]interface{}, 0),
				}

				for _, item := range shopOrder.Items {
					itemBrief := new(ItemBrief)
					if err := mapper.Map(item, itemBrief); err != nil {
						return err
					}

					// search for promotion id to apply
					for _, model := range item.Models {
						if itemBrief.Modelid == model.Modelid {
							itemBrief.AppliedPromotionID = model.Promotionid

							break
						}
					}

					// assign cart item change unixtime
					itemBrief.CartItemChangeTime = shopOrder.Shop.AddinTime

					newShopOrder.ItemBriefs = append(newShopOrder.ItemBriefs, *itemBrief)
				}

				shopOrders = append(shopOrders, newShopOrder)
			}
		}
	}

	body["selected_shop_order_ids"] = shopOrders
	body["platform_vouchers"] = make([]interface{}, 0)

	raw, err := cart.sh.post("/v4/cart/checkout", "https://shopee.co.id/cart/", body)
	if err != nil {
		return err
	}

	info := new(Cart)
	if err := json.Unmarshal(raw, info); err != nil {
		return err
	}

	if info.Error > 0 {
		return errors.New(info.ErrorMessage)
	}

	cart.SelectedShopOrderIDS = shopOrders

	return nil
}
