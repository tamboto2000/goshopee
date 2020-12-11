package goshopee

import (
	"encoding/json"
	"errors"

	"github.com/tamboto2000/mapper"
	"github.com/tamboto2000/random"
)

// Cart actions
const (
	SelectAllItems   = 4
	DeselectAllItems = 5
)

// Cart contains info about your cart
type Cart struct {
	Data                       *Data                 `json:"data,omitempty"`
	Error                      int                   `json:"error,omitempty"`
	ErrorMessage               string                `json:"error_message,omitempty"`
	MessageLevel               *MessageLevel         `json:"message_level,omitempty"`
	WarnMessage                interface{}           `json:"warn_message,omitempty"`
	SelectedShopOrderIDS       []ShopOrder           `json:"selected_shop_order_ids,omitempty"`
	Shoporders                 []ShopOrder           `json:"shoporders,omitempty"`
	PlatformVouchers           []interface{}         `json:"platform_vouchers,omitempty"`
	Addresses                  *UserAddresses        `json:"addresses,omitempty"`
	SelectedPaymentChannelData *PaymentChannelData   `json:"selected_payment_channel_data,omitempty"`
	PromotionData              *PromotionData        `json:"promotion_data,omitempty"`
	DeviceInfo                 *DeviceInfo           `json:"device_info,omitempty"`
	TaxInfo                    *TaxInfo              `json:"tax_info,omitempty"`
	CartType                   int                   `json:"cart_type,omitempty"`
	ShippingOrders             []ShippingOrder       `json:"shipping_orders,omitempty"`
	DisabledCheckoutInfo       *DisabledCheckoutInfo `json:"disabled_checkout_info,omitempty"`
	Timestamp                  int                   `json:"timestamp,omitempty"`
	CheckoutPriceData          *PriceData            `json:"checkout_price_data,omitempty"`
	ClientID                   int                   `json:"client_id,omitempty"`
	PaymentChannelInfo         *PaymentChannelInfo   `json:"payment_channel_info,omitempty"`
	DropshippingInfo           *DropshippingInfo     `json:"dropshipping_info,omitempty"`
	CanCheckout                bool                  `json:"can_checkout,omitempty"`
	OrderUpdateInfo            *OrderUpdateInfo      `json:"order_update_info,omitempty"`
	BuyerTxnFeeInfo            *BuyerTxnFeeInfo      `json:"buyer_txn_fee_info,omitempty"`
	// This field contains a prepared request data for place order after calling CheckoutAll. Actually the Cart itself is
	// can be used for request body, but there's a lot of unecessary data and can slow down place order process,
	// but nontheless the Cart will be updated after calling CheckoutAll method
	PrepCheckout *Cart `json:"prepCheckout,omitempty"`
	// This field contains information about placed order after calling PlaceOrder
	PlacedOrder *Order `json:"placedOrder,omitempty"`
	Status      int    `json:"status,omitempty"`

	sh *Shopee
}

type PriceData struct {
	ShippingSubtotal               int `json:"shipping_subtotal,omitempty"`
	ShippingDiscountSubtotal       int `json:"shipping_discount_subtotal,omitempty"`
	ShippingSubtotalBeforeDiscount int `json:"shipping_subtotal_before_discount,omitempty"`
	BundleDealsDiscount            int `json:"bundle_deals_discount,omitempty"`
	GroupBuyDiscount               int `json:"group_buy_discount,omitempty"`
	MerchandiseSubtotal            int `json:"merchandise_subtotal,omitempty"`
	TaxPayable                     int `json:"tax_payable,omitempty"`
	BuyerTxnFee                    int `json:"buyer_txn_fee,omitempty"`
	CreditCardPromotion            int `json:"credit_card_promotion,omitempty"`
	PromocodeApplied               int `json:"promocode_applied,omitempty"`
	ShopeeCoinsRedeemed            int `json:"shopee_coins_redeemed,omitempty"`
	TotalPayable                   int `json:"total_payable,omitempty"`
}

type BuyerTxnFeeInfo struct {
	Error string `json:"error,omitempty"`
}

type DisabledCheckoutInfo struct {
	AutoPopup   bool        `json:"auto_popup,omitempty"`
	Description string      `json:"description,omitempty"`
	ErrorInfos  []ErrorInfo `json:"error_infos,omitempty"`
}

type ErrorInfo struct {
	Message     interface{} `json:"message,omitempty"`
	ErrorType   string      `json:"error_type,omitempty"`
	ShippingID  int         `json:"shipping_id,omitempty"`
	ErrorAction string      `json:"error_action,omitempty"`
}

type DropshippingInfo struct {
	PhoneNumber string `json:"phone_number,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Name        string `json:"name,omitempty"`
}

type OrderUpdateInfo struct {
}

type PaymentChannelInfo struct {
	Channels      []Channel             `json:"channels,omitempty"`
	PromotionInfo *ChannelPromotionInfo `json:"promotion_info,omitempty"`
	GroupingInfo  *GroupingInfo         `json:"grouping_info,omitempty"`
}

type Channel struct {
	PreselectDisabled        bool                     `json:"preselect_disabled,omitempty"`
	Description              string                   `json:"description,omitempty"`
	InfoLink                 string                   `json:"info_link,omitempty"`
	IsNew                    bool                     `json:"is_new,omitempty"`
	PromotionInfo            *ChannelPromotionInfo    `json:"promotion_info,omitempty"`
	Groups                   *Groups                  `json:"groups,omitempty"`
	Icon                     string                   `json:"icon,omitempty"`
	Name                     string                   `json:"name,omitempty"`
	InfoText                 string                   `json:"info_text,omitempty"`
	PaymentResultText        string                   `json:"payment_result_text,omitempty"`
	Enabled                  bool                     `json:"enabled,omitempty"`
	ChannelID                int                      `json:"channel_id,omitempty"`
	Priority                 int                      `json:"priority,omitempty"`
	IconBackground           interface{}              `json:"icon_background,omitempty"`
	Version                  int                      `json:"version,omitempty"`
	PaymentChannelHints      *OrderUpdateInfo         `json:"payment_channel_hints,omitempty"`
	ChannelBehavior          *ChannelBehavior         `json:"channel_behavior,omitempty"`
	Balance                  int                      `json:"balance,omitempty"`
	SupportSubscription      bool                     `json:"support_subscription,omitempty"`
	Banks                    []Bank                   `json:"banks,omitempty"`
	Promotions               []Promotion              `json:"promotions,omitempty"`
	AddCardURL               *AddCardURL              `json:"add_card_url,omitempty"`
	Cards                    []interface{}            `json:"cards,omitempty"`
	DisabledReasonKey        int                      `json:"disabled_reason_key,omitempty"`
	Subcategory              int                      `json:"subcategory,omitempty"`
	Currency                 string                   `json:"currency,omitempty"`
	Flag                     int                      `json:"flag,omitempty"`
	DisabledReasonArgs       *DisabledReasonArgs      `json:"disabled_reason_args,omitempty"`
	SPMOptionInfo            string                   `json:"spm_option_info,omitempty"`
	Category                 int                      `json:"category,omitempty"`
	ExtraData                *ChannelExtraData        `json:"extra_data,omitempty"`
	Country                  string                   `json:"country,omitempty"`
	Channelid                int                      `json:"channelid,omitempty"`
	NameLabel                string                   `json:"name_label,omitempty"`
	DisabledReason           string                   `json:"disabled_reason,omitempty"`
	SPMChannelID             int                      `json:"spm_channel_id,omitempty"`
	InstallmentBanks         []InstallmentBank        `json:"installment_banks,omitempty"`
	IsCreditCardChannel      int                      `json:"is_credit_card_channel,omitempty"`
	HasCards                 bool                     `json:"has_cards,omitempty"`
	CreditCards              []interface{}            `json:"credit_cards,omitempty"`
	InstallmentPromotionLink string                   `json:"installment_promotion_link,omitempty"`
	InstallmentPlans         []ChannelInstallmentPlan `json:"installment_plans,omitempty"`
	InstallmentData          *InstallmentData         `json:"installment_data,omitempty"`
	InstallmentPromotionText string                   `json:"installment_promotion_text,omitempty"`
	Subchannels              []Subchannel             `json:"subchannels,omitempty"`
}

type AddCardURL struct {
	API string `json:"api,omitempty"`
	UI  string `json:"ui,omitempty"`
}

type DisabledReasonData struct {
	ID               int    `json:"id,omitempty"`
	MessageLocalized string `json:"message_localized,omitempty"`
	Description      string `json:"description,omitempty"`
	MinAmount        int    `json:"min_amount,omitempty"`
}

type PopupConfirmationData struct {
	CancelButtonMessage string `json:"cancel_button_message,omitempty"`
	NeedPopup           bool   `json:"need_popup,omitempty"`
	Message             string `json:"message,omitempty"`
	OkButtonMessage     string `json:"ok_button_message,omitempty"`
}

type BankSubDescriptionInfo struct {
	Important string `json:"important,omitempty"`
	Normal    string `json:"normal,omitempty"`
}

type ChannelBehavior struct {
	DisableInstruction bool `json:"disable_instruction,omitempty"`
	DisableCancel      bool `json:"disable_cancel,omitempty"`
}

type DisabledReasonArgs struct {
	ChannelName string `json:"channel_name,omitempty"`
}

type ChannelExtraData struct {
	VoucherPaymentType int    `json:"voucher_payment_type,omitempty"`
	BannedCategories   []int  `json:"banned_categories,omitempty"`
	ExpiryDuration     int    `json:"expiry_duration,omitempty"`
	SPMChannelID       int    `json:"spm_channel_id,omitempty"`
	RedirectURL        string `json:"redirect_url,omitempty"`
	PriceLimit         int    `json:"price_limit,omitempty"`
	ExpiryExtension    int    `json:"expiry_extension,omitempty"`
}

type Groups struct {
	BankTransfer bool `json:"bank_transfer,omitempty"`
	Installment  bool `json:"installment,omitempty"`
	Immediate    bool `json:"immediate,omitempty"`
	CreditCard   bool `json:"credit_card,omitempty"`
	Wallet       bool `json:"wallet,omitempty"`
	BankAccount  bool `json:"bank_account,omitempty"`
}

type InstallmentBank struct {
	BankName         string                           `json:"bank_name,omitempty"`
	Subdescription   string                           `json:"subdescription,omitempty"`
	Enabled          bool                             `json:"enabled,omitempty"`
	BankID           int                              `json:"bank_id,omitempty"`
	Priority         int                              `json:"priority,omitempty"`
	IconBackground   interface{}                      `json:"icon_background,omitempty"`
	Cards            []interface{}                    `json:"cards,omitempty"`
	InstallmentPlans []InstallmentBankInstallmentPlan `json:"installment_plans,omitempty"`
	DisabledReason   string                           `json:"disabled_reason,omitempty"`
	Icon             string                           `json:"icon,omitempty"`
}

type InstallmentBankInstallmentPlan struct {
	DisabledReasonKey  string `json:"disabled_reason_key,omitempty"`
	DisabledReasonData *Data  `json:"disabled_reason_data,omitempty"`
	MonthlyPayment     int    `json:"monthly_payment,omitempty"`
	OptionInfo         string `json:"option_info,omitempty"`
	DisabledReason     string `json:"disabled_reason,omitempty"`
	Enabled            bool   `json:"enabled,omitempty"`
	BankID             int    `json:"bank_id,omitempty"`
	Tenure             int    `json:"tenure,omitempty"`
	PlanName           string `json:"plan_name,omitempty"`
	Description        string `json:"description,omitempty"`
}

type InstallmentData struct {
	PlanInfo     string `json:"plan_info,omitempty"`
	PlanInfoLink string `json:"plan_info_link,omitempty"`
}

type ChannelInstallmentPlan struct {
	MonthlyPayment     int              `json:"monthly_payment,omitempty"`
	Description        string           `json:"description,omitempty"`
	Enabled            bool             `json:"enabled,omitempty"`
	IconBackground     interface{}      `json:"icon_background,omitempty"`
	Tenure             int              `json:"tenure,omitempty"`
	Icon               string           `json:"icon,omitempty"`
	PlanName           string           `json:"plan_name,omitempty"`
	OptionInfo         string           `json:"option_info,omitempty"`
	DisabledReasonKey  string           `json:"disabled_reason_key,omitempty"`
	DisabledReasonData *OrderUpdateInfo `json:"disabled_reason_data,omitempty"`
	DisabledReason     string           `json:"disabled_reason,omitempty"`
}

type ChannelPromotionInfo struct {
	Text string `json:"text,omitempty"`
}

type Promotion struct {
	Description     string `json:"description,omitempty"`
	Title           string `json:"title,omitempty"`
	URL             string `json:"url,omitempty"`
	BankID          int    `json:"bank_id,omitempty"`
	DiscountType    int    `json:"discount_type,omitempty"`
	BankLogo        string `json:"bank_logo,omitempty"`
	DiscountValue   int    `json:"discount_value,omitempty"`
	CardPromotionID int    `json:"card_promotion_id,omitempty"`
	PrimaryColor    string `json:"primary_color,omitempty"`
}

type Subchannel struct {
	Subcategory         int                           `json:"subcategory,omitempty"`
	Enabled             bool                          `json:"enabled,omitempty"`
	ParentChannelid     int                           `json:"parent_channelid,omitempty"`
	PromotionInfo       *ChannelPromotionInfo         `json:"promotion_info,omitempty"`
	Currency            string                        `json:"currency,omitempty"`
	Flag                int                           `json:"flag,omitempty"`
	SPMOptionInfo       string                        `json:"spm_option_info,omitempty"`
	Description         string                        `json:"description,omitempty"`
	Category            int                           `json:"category,omitempty"`
	ExtraData           *SubchannelExtraData          `json:"extra_data,omitempty"`
	Name                string                        `json:"name,omitempty"`
	Country             string                        `json:"country,omitempty"`
	Channelid           int                           `json:"channelid,omitempty"`
	NameLabel           string                        `json:"name_label,omitempty"`
	Priority            int                           `json:"priority,omitempty"`
	IconBackground      interface{}                   `json:"icon_background,omitempty"`
	Version             int                           `json:"version,omitempty"`
	SubDescriptionInfo  *SubchannelSubDescriptionInfo `json:"sub_description_info,omitempty"`
	PaymentChannelHints *OrderUpdateInfo              `json:"payment_channel_hints,omitempty"`
	Icon                string                        `json:"icon,omitempty"`
	SPMChannelID        int                           `json:"spm_channel_id,omitempty"`
}

type SubchannelExtraData struct {
	EligibleTransfer string `json:"eligible_transfer,omitempty"`
}

type SubchannelSubDescriptionInfo struct {
	Second *First `json:"second,omitempty"`
	First  *First `json:"first,omitempty"`
}

type First struct {
	Text      string `json:"text,omitempty"`
	Highlight bool   `json:"highlight,omitempty"`
}

type GroupingInfo struct {
	Groups []Group `json:"groups,omitempty"`
}

type Group struct {
	DisplayInfo *DisplayInfo `json:"display_info,omitempty"`
	GroupID     interface{}  `json:"group_id,omitempty"`
	ToCombine   []int        `json:"to_combine,omitempty"`
}

type DisplayInfo struct {
	Name           string                    `json:"name,omitempty"`
	InfoLink       string                    `json:"info_link,omitempty"`
	InfoText       string                    `json:"info_text,omitempty"`
	IsNew          bool                      `json:"is_new,omitempty"`
	PromotionInfo  *DisplayInfoPromotionInfo `json:"promotion_info,omitempty"`
	IconBackground string                    `json:"icon_background,omitempty"`
	ChannelID      int                       `json:"channel_id,omitempty"`
	Version        int                       `json:"version,omitempty"`
	Icon           string                    `json:"icon,omitempty"`
}

type DisplayInfoPromotionInfo struct {
	VoucherPaymentType interface{} `json:"voucher_payment_type,omitempty"`
}

type ShopVoucherEntrance struct {
	Status bool `json:"status,omitempty"`
	Shopid int  `json:"shopid,omitempty"`
}

type VoucherInfo struct {
	CoinEarned         int         `json:"coin_earned,omitempty"`
	Promotionid        int         `json:"promotionid,omitempty"`
	DiscountPercentage int         `json:"discount_percentage,omitempty"`
	DiscountValue      int         `json:"discount_value,omitempty"`
	VoucherCode        interface{} `json:"voucher_code,omitempty"`
	RewardType         int         `json:"reward_type,omitempty"`
	CoinPercentage     int         `json:"coin_percentage,omitempty"`
	UsedPrice          int         `json:"used_price,omitempty"`
}

type SelectedPaymentChannelData struct {
	ChannelID             int                   `json:"channel_id,omitempty"`
	ChannelItemOptionInfo ChannelItemOptionInfo `json:"channel_item_option_info,omitempty"`
	Version               int                   `json:"version,omitempty"`
}

type ChannelItemOptionInfo struct {
	OptionInfo string `json:"option_info,omitempty"`
}

type ShippingOrder struct {
	BuyerRemark                          interface{}       `json:"buyer_remark,omitempty"`
	Logistics                            *Logistics        `json:"logistics,omitempty"`
	TaxPayable                           int               `json:"tax_payable,omitempty"`
	OrderTotal                           int               `json:"order_total,omitempty"`
	ShippingID                           int               `json:"shipping_id,omitempty"`
	BuyerICNumber                        interface{}       `json:"buyer_ic_number,omitempty"`
	FulfillmentInfo                      *FulfillmentInfo  `json:"fulfillment_info,omitempty"`
	ShopeeShippingDiscountID             int               `json:"shopee_shipping_discount_id,omitempty"`
	SelectedLogisticChannelidWithWarning interface{}       `json:"selected_logistic_channelid_with_warning,omitempty"`
	SelectedLogisticChannelid            int               `json:"selected_logistic_channelid,omitempty"`
	VoucherWalletCheckingChannelIDS      []interface{}     `json:"voucher_wallet_checking_channel_ids,omitempty"`
	CodFee                               int               `json:"cod_fee,omitempty"`
	ShoporderIndexes                     []int             `json:"shoporder_indexes,omitempty"`
	BuyerAddressData                     *BuyerAddressData `json:"buyer_address_data,omitempty"`
	ShippingFeeDiscount                  int               `json:"shipping_fee_discount,omitempty"`
	ShippingGroupDescription             string            `json:"shipping_group_description,omitempty"`
	AmountDetail                         *Detail           `json:"amount_detail,omitempty"`
	OrderTotalWithoutShipping            int               `json:"order_total_without_shipping,omitempty"`
	ShippingFee                          int               `json:"shipping_fee,omitempty"`
	ShippingGroupIcon                    string            `json:"shipping_group_icon,omitempty"`
}

type Detail struct {
	BasicShippingFee                int  `json:"BASIC_SHIPPING_FEE,omitempty"`
	SellerEstimatedInsuranceFee     int  `json:"SELLER_ESTIMATED_INSURANCE_FEE,omitempty"`
	ShopeeOrSellerShippingDiscount  int  `json:"SHOPEE_OR_SELLER_SHIPPING_DISCOUNT,omitempty"`
	VoucherDiscount                 int  `json:"VOUCHER_DISCOUNT,omitempty"`
	ShippingDiscountBySeller        int  `json:"SHIPPING_DISCOUNT_BY_SELLER,omitempty"`
	SellerEstimatedBasicShippingFee int  `json:"SELLER_ESTIMATED_BASIC_SHIPPING_FEE,omitempty"`
	ShippingDiscountByShopee        int  `json:"SHIPPING_DISCOUNT_BY_SHOPEE,omitempty"`
	InsuranceFee                    int  `json:"INSURANCE_FEE,omitempty"`
	ItemTotal                       int  `json:"ITEM_TOTAL,omitempty"`
	ShopPromoOnly                   bool `json:"shop_promo_only,omitempty"`
	CodFee                          int  `json:"COD_FEE,omitempty"`
	TaxFee                          int  `json:"TAX_FEE,omitempty"`
	SellerOnlyShippingDiscount      int  `json:"SELLER_ONLY_SHIPPING_DISCOUNT,omitempty"`
}

type BuyerAddressData struct {
	TaxAddress  string `json:"tax_address,omitempty"`
	ErrorStatus string `json:"error_status,omitempty"`
	AddressType int    `json:"address_type,omitempty"`
	Addressid   int    `json:"addressid,omitempty"`
}

type FulfillmentInfo struct {
	FulfillmentFlag      int    `json:"fulfillment_flag,omitempty"`
	FulfillmentSource    string `json:"fulfillment_source,omitempty"`
	ManagedBySbs         bool   `json:"managed_by_sbs,omitempty"`
	OrderFulfillmentType int    `json:"order_fulfillment_type,omitempty"`
	WarehouseAddressID   int    `json:"warehouse_address_id,omitempty"`
}

type LogisticChannel struct {
	ChannelData                   *ChannelData                   `json:"channel_data,omitempty"`
	CostInfo                      *CostInfo                      `json:"cost_info,omitempty"`
	PreferredDeliveryInstructions *PreferredDeliveryInstructions `json:"preferred_delivery_instructions,omitempty"`
	ShippabilityInfo              *ShippabilityInfo              `json:"shippability_info,omitempty"`
	ShippingFeeData               *ShippingFeeData               `json:"shipping_fee_data,omitempty"`
	DaysToDeliver                 int                            `json:"days_to_deliver,omitempty"`
	PreferredDeliveryTimeInfo     *PreferredDeliveryTimeInfo     `json:"preferred_delivery_time_info,omitempty"`
	CoverShippingFee              bool                           `json:"cover_shipping_fee,omitempty"`
	DeliveryData                  *DeliveryData                  `json:"delivery_data,omitempty"`
	CodData                       *CodData                       `json:"cod_data,omitempty"`
	GuaranteeExtensionPeriod      interface{}                    `json:"guarantee_extension_period,omitempty"`
}

type ChannelData struct {
	Category                      int     `json:"category,omitempty"`
	Cashless                      int     `json:"cashless,omitempty"`
	DisplayName                   string  `json:"display_name,omitempty"`
	Name                          string  `json:"name,omitempty"`
	IsMaskChannel                 int     `json:"is_mask_channel,omitempty"`
	Channelid                     int     `json:"channelid,omitempty"`
	Enabled                       bool    `json:"enabled,omitempty"`
	ExtraFlag                     int     `json:"extra_flag,omitempty"`
	MinAmountNeedIC               int     `json:"min_amount_need_ic,omitempty"`
	Priority                      int     `json:"priority,omitempty"`
	WarningMsg                    string  `json:"warning_msg,omitempty"`
	Flag                          float64 `json:"flag,omitempty"`
	Warning                       string  `json:"warning,omitempty"`
	Maintenance                   int     `json:"maintenance,omitempty"`
	NeedCheckIC                   bool    `json:"need_check_ic,omitempty"`
	IsShowPreferredDeliveryOption int     `json:"is_show_preferred_delivery_option,omitempty"`
	AddressType                   int     `json:"address_type,omitempty"`
	CodSupported                  bool    `json:"cod_supported,omitempty"`
	ShippingMethod                int     `json:"shipping_method,omitempty"`
}

type CodData struct {
	CodAvailable bool `json:"cod_available,omitempty"`
}

type CostInfo struct {
	DiscountPromotionRuleSnapshotID int        `json:"discount_promotion_rule_snapshot_id,omitempty"`
	EstimatedShippingFee            int        `json:"estimated_shipping_fee,omitempty"`
	EnjoyDiscount                   bool       `json:"enjoy_discount,omitempty"`
	DiscountedShippingFee           int        `json:"discounted_shipping_fee,omitempty"`
	DiscountPromotionRuleID         int        `json:"discount_promotion_rule_id,omitempty"`
	RebatePromotionRuleID           int        `json:"rebate_promotion_rule_id,omitempty"`
	Discounts                       *Discounts `json:"discounts,omitempty"`
	Cost                            int        `json:"cost,omitempty"`
	OriginalCost                    int        `json:"original_cost,omitempty"`
	DiscountAmount                  int        `json:"discount_amount,omitempty"`
}

type Discounts struct {
	Shopee int `json:"shopee,omitempty"`
	Seller int `json:"seller,omitempty"`
}

type DeliveryData struct {
	DelayMessage              string      `json:"delay_message,omitempty"`
	EstimatedDeliveryDateFrom int         `json:"estimated_delivery_date_from,omitempty"`
	IsShopee24H               bool        `json:"is_shopee_24h,omitempty"`
	MinDays                   int         `json:"min_days,omitempty"`
	DetailInfo                *DetailInfo `json:"detail_info,omitempty"`
	EstimatedDeliveryDateTo   int         `json:"estimated_delivery_date_to,omitempty"`
	IsRapidSla                bool        `json:"is_rapid_sla,omitempty"`
	MaxDays                   int         `json:"max_days,omitempty"`
	EstimatedDeliveryTimeMin  int         `json:"estimated_delivery_time_min,omitempty"`
	DisplayMode               string      `json:"display_mode,omitempty"`
	IsCrossBorder             bool        `json:"is_cross_border,omitempty"`
	EstimatedDeliveryTimeMax  int         `json:"estimated_delivery_time_max,omitempty"`
	HasEdt                    bool        `json:"has_edt,omitempty"`
}

type DetailInfo struct {
	EdtMaxDt string  `json:"edt_max_dt,omitempty"`
	HePt     int     `json:"he_pt,omitempty"`
	CdtMax   float64 `json:"cdt_max,omitempty"`
	HeCdt    int     `json:"he_cdt,omitempty"`
	Apt      float64 `json:"apt,omitempty"`
	CdtMin   float64 `json:"cdt_min,omitempty"`
	EdtMinDt string  `json:"edt_min_dt,omitempty"`
}

type PreferredDeliveryInstructions struct {
	Options []interface{} `json:"options,omitempty"`
}

type PreferredDeliveryTimeInfo struct {
	WarningMsg string   `json:"warning_msg,omitempty"`
	AlertMsg   string   `json:"alert_msg,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

type ShippabilityInfo struct {
	NextAvailableDate string `json:"next_available_date,omitempty"`
	AbleToShip        bool   `json:"able_to_ship,omitempty"`
}

type ShippingFeeData struct {
	CodFee                    int     `json:"cod_fee,omitempty"`
	ShippingFeeBeforeDiscount int     `json:"shipping_fee_before_discount,omitempty"`
	ShippingFeeDetail         *Detail `json:"shipping_fee_detail,omitempty"`
	ChargeableShippingFee     int     `json:"chargeable_shipping_fee,omitempty"`
}

type LogisticServiceTypes struct {
	SelfCollection Regular `json:"self_collection,omitempty"`
	Regular        Regular `json:"regular,omitempty"`
	RegularCargo   Regular `json:"regular_cargo,omitempty"`
}

type Regular struct {
	Name                     string                   `json:"name,omitempty"`
	MaxCost                  int                      `json:"max_cost,omitempty"`
	ChannelIDS               []int                    `json:"channel_ids,omitempty"`
	MinCost                  int                      `json:"min_cost,omitempty"`
	Priority                 int                      `json:"priority,omitempty"`
	ConsolidatedLogisticInfo ConsolidatedLogisticInfo `json:"consolidated_logistic_info,omitempty"`
	Identifier               string                   `json:"identifier,omitempty"`
	Enabled                  bool                     `json:"enabled,omitempty"`
	SlaMsg                   string                   `json:"sla_msg,omitempty"`
}

type ConsolidatedLogisticInfo struct {
	ShopeeMaxCost        int  `json:"shopee_max_cost,omitempty"`
	ShopPromoOnlyMinCost int  `json:"shop_promo_only_min_cost,omitempty"`
	ShopPromoOnlyMaxCost int  `json:"shop_promo_only_max_cost,omitempty"`
	Enabled              bool `json:"enabled,omitempty"`
	ShopeeMinCost        int  `json:"shopee_min_cost,omitempty"`
}

type Shoporder struct {
	Shop                      *Shop             `json:"shop,omitempty"`
	BuyerRemark               interface{}       `json:"buyer_remark,omitempty"`
	ShippingFee               int               `json:"shipping_fee,omitempty"`
	OrderTotal                int               `json:"order_total,omitempty"`
	ShippingID                int               `json:"shipping_id,omitempty"`
	BuyerICNumber             interface{}       `json:"buyer_ic_number,omitempty"`
	Items                     []Item            `json:"items,omitempty"`
	Logistics                 *Logistics        `json:"logistics,omitempty"`
	SelectedLogisticChannelid int               `json:"selected_logistic_channelid,omitempty"`
	CodFee                    int               `json:"cod_fee,omitempty"`
	TaxPayable                int               `json:"tax_payable,omitempty"`
	BuyerAddressData          *BuyerAddressData `json:"buyer_address_data,omitempty"`
	ShippingFeeDiscount       int               `json:"shipping_fee_discount,omitempty"`
	OrderTotalWithoutShipping int               `json:"order_total_without_shipping,omitempty"`
	AmountDetail              *Detail           `json:"amount_detail,omitempty"`
}

type OpcExtraData struct {
	SlashPriceActivityID int `json:"slash_price_activity_id,omitempty"`
}

type PromotionData struct {
	UseCoins                  bool                     `json:"use_coins,omitempty"`
	FreeShippingVoucherInfo   *FreeShippingVoucherInfo `json:"free_shipping_voucher_info,omitempty"`
	PlatformVouchers          []interface{}            `json:"platform_vouchers,omitempty"`
	ShopVouchers              []interface{}            `json:"shop_vouchers,omitempty"`
	CheckShopVoucherEntrances bool                     `json:"check_shop_voucher_entrances,omitempty"`
	AutoApplyShopVoucher      bool                     `json:"auto_apply_shop_voucher,omitempty"`
	PromotionMsg              string                   `json:"promotion_msg,omitempty"`
	PriceDiscount             int                      `json:"price_discount,omitempty"`
	CanUseCoins               bool                     `json:"can_use_coins,omitempty"`
	VoucherInfo               *VoucherInfo             `json:"voucher_info,omitempty"`
	CoinInfo                  *CoinInfo                `json:"coin_info,omitempty"`
	AppliedVoucherCode        interface{}              `json:"applied_voucher_code,omitempty"`
	ShopVoucherEntrances      []ShopVoucherEntrance    `json:"shop_voucher_entrances,omitempty"`
	CardPromotionEnabled      bool                     `json:"card_promotion_enabled,omitempty"`
	InvalidMessage            interface{}              `json:"invalid_message,omitempty"`
	CardPromotionID           interface{}              `json:"card_promotion_id,omitempty"`
	VoucherCode               interface{}              `json:"voucher_code,omitempty"`
}

type DeviceInfo struct {
	DeviceID          string              `json:"device_id,omitempty"`
	DeviceFingerprint string              `json:"device_fingerprint,omitempty"`
	TongdunBlackbox   string              `json:"tongdun_blackbox,omitempty"`
	BuyerPaymentInfo  *PaymentChannelData `json:"buyer_payment_info,omitempty"`
}

type TaxInfo struct {
	TaxID string `json:"tax_id,omitempty"`
}

type MessageLevel struct {
	Toast   bool        `json:"toast,omitempty"`
	Popup   bool        `json:"popup,omitempty"`
	Refresh interface{} `json:"refresh,omitempty"`
}

type Order struct {
	IsSlashPrice        bool            `json:"is_slash_price,omitempty"`
	IsShopeeCod         bool            `json:"is_shopee_cod,omitempty"`
	Timestamp           int             `json:"timestamp,omitempty"`
	NavigationInfo      *NavigationInfo `json:"navigation_info,omitempty"`
	GroupBuyGroupid     int             `json:"group_buy_groupid,omitempty"`
	Checkoutid          int             `json:"checkoutid,omitempty"`
	HasNonShippableItem bool            `json:"has_non_shippable_item,omitempty"`
	RedirectURL         string          `json:"redirect_url,omitempty"`
	Orderids            []int           `json:"orderids,omitempty"`
	IsGroupBuy          bool            `json:"is_group_buy,omitempty"`
	PaymentType         int             `json:"payment_type,omitempty"`
}

type NavigationInfo struct {
	Data Data `json:"data,omitempty"`
	Type int  `json:"type,omitempty"`
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

// SyncAddresses get user's addresses, assigned to Cart
func (cart *Cart) SyncAddresses() error {
	raw, err := cart.sh.get("/v1/addresses/", nil)
	if err != nil {
		return err
	}

	userAddr := new(UserAddresses)
	if err := json.Unmarshal(raw, userAddr); err != nil {
		return err
	}

	cart.Addresses = userAddr

	return nil
}

// CheckoutAll checkout all items in cart.
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
	if err := cart.detailedCheckout(); err != nil {
		return err
	}

	return nil
}

// PlaceOrder place your order to queue.
// May God and The HTTP Overlord bless this function...
func (cart *Cart) PlaceOrder() error {
	// raw, err := cart.sh.post("/v2/checkout/place_order", "https://shopee.co.id/checkout/", cart.PrepCheckout)
	raw, err := cart.sh.customPost(baseAPIURL+"/v2/checkout/place_order", "https://shopee.co.id/checkout/", map[string]string{
		"X-CV-ID":    "7",
		"X-Track-Id": random.RandHexStr(128),
	}, cart.PrepCheckout)
	if err != nil {
		return err
	}

	order := new(Order)
	if err := json.Unmarshal(raw, order); err != nil {
		return err
	}

	cart.PlacedOrder = order

	return nil
}

func (cart *Cart) detailedCheckout() error {
	body := make(map[string]interface{})
	newShopOrders := make([]ShopOrder, 0)
	for _, shopOrder := range cart.SelectedShopOrderIDS {
		newShopOrder := ShopOrder{Shop: &Shop{Shopid: shopOrder.Shopid}}

		for _, itemBrief := range shopOrder.ItemBriefs {
			item := new(Item)
			if err := mapper.Map(itemBrief, item); err != nil {
				return err
			}

			newShopOrder.Items = append(newShopOrder.Items, *item)
		}

		newShopOrders = append(newShopOrders, newShopOrder)
	}

	body["shoporders"] = newShopOrders
	raw, err := cart.sh.post("/v2/checkout/get", "https://shopee.co.id/checkout", body)
	if err != nil {
		return err
	}

	newCart := new(Cart)
	if err := json.Unmarshal(raw, newCart); err != nil {
		return err
	}

	if newCart.Error > 0 {
		return errors.New(newCart.ErrorMessage)
	}

	newCart.Status = 200
	newCart.ClientID = 0
	cart.PrepCheckout = newCart

	return mapper.MapWithOpt(newCart, cart, mapper.FieldOption{SkipAssigned: true})
}

// Encrypt Shopeepay pin.
// May God bless this function...
func (cart *Cart) encryptPin() (string, error) {

	return "", nil
}
