package goshopee

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// Item contains item or product detail
type Item struct {
	Item                                  *Item             `json:"item,omitempty"`
	Itemid                                int               `json:"itemid,omitempty"`
	PriceMaxBeforeDiscount                int               `json:"price_max_before_discount,omitempty"`
	ItemStatus                            string            `json:"item_status,omitempty"`
	CanUseWholesale                       bool              `json:"can_use_wholesale,omitempty"`
	ShowFreeShipping                      bool              `json:"show_free_shipping,omitempty"`
	EstimatedDays                         int               `json:"estimated_days,omitempty"`
	IsHotSales                            bool              `json:"is_hot_sales,omitempty"`
	IsSlashPriceItem                      bool              `json:"is_slash_price_item,omitempty"`
	UpcomingFlashSale                     *Item             `json:"upcoming_flash_sale,omitempty"`
	SlashLowestPrice                      interface{}       `json:"slash_lowest_price,omitempty"`
	IsPartialFulfilled                    bool              `json:"is_partial_fulfilled,omitempty"`
	Condition                             int               `json:"condition,omitempty"`
	IsNonCcInstallmentPaymentEligible     bool              `json:"is_non_cc_installment_payment_eligible,omitempty"`
	Categories                            []Category        `json:"categories,omitempty"`
	Ctime                                 int               `json:"ctime,omitempty"`
	Name                                  string            `json:"name,omitempty"`
	ShowShopeeVerifiedLabel               bool              `json:"show_shopee_verified_label,omitempty"`
	SizeChart                             interface{}       `json:"size_chart,omitempty"`
	IsPreOrder                            bool              `json:"is_pre_order,omitempty"`
	ServiceByShopeeFlag                   interface{}       `json:"service_by_shopee_flag,omitempty"`
	HistoricalSold                        int               `json:"historical_sold,omitempty"`
	ReferenceItemID                       string            `json:"reference_item_id,omitempty"`
	RecommendationInfo                    interface{}       `json:"recommendation_info,omitempty"`
	BundleDealInfo                        interface{}       `json:"bundle_deal_info,omitempty"`
	PriceMax                              int               `json:"price_max,omitempty"`
	HasLowestPriceGuarantee               bool              `json:"has_lowest_price_guarantee,omitempty"`
	ShippingIconType                      int               `json:"shipping_icon_type,omitempty"`
	Images                                []string          `json:"images,omitempty"`
	PriceBeforeDiscount                   int               `json:"price_before_discount,omitempty"`
	CodFlag                               int               `json:"cod_flag,omitempty"`
	Catid                                 int               `json:"catid,omitempty"`
	IsOfficialShop                        bool              `json:"is_official_shop,omitempty"`
	CoinEarnLabel                         interface{}       `json:"coin_earn_label,omitempty"`
	HashtagList                           interface{}       `json:"hashtag_list,omitempty"`
	Sold                                  int               `json:"sold,omitempty"`
	Makeup                                interface{}       `json:"makeup,omitempty"`
	ItemRating                            *Rating           `json:"item_rating,omitempty"`
	ShowOfficialShopLabelInTitle          bool              `json:"show_official_shop_label_in_title,omitempty"`
	Discount                              string            `json:"discount,omitempty"`
	Reason                                interface{}       `json:"reason,omitempty"`
	LabelIDS                              []int             `json:"label_ids,omitempty"`
	HasGroupBuyStock                      bool              `json:"has_group_buy_stock,omitempty"`
	OtherStock                            interface{}       `json:"other_stock,omitempty"`
	DeepDiscount                          interface{}       `json:"deep_discount,omitempty"`
	Attributes                            []Attribute       `json:"attributes,omitempty"`
	BadgeIconType                         int               `json:"badge_icon_type,omitempty"`
	Liked                                 bool              `json:"liked,omitempty"`
	CmtCount                              int               `json:"cmt_count,omitempty"`
	Image                                 string            `json:"image,omitempty"`
	RecommendationAlgorithm               interface{}       `json:"recommendation_algorithm,omitempty"`
	IsCcInstallmentPaymentEligible        bool              `json:"is_cc_installment_payment_eligible,omitempty"`
	Shopid                                int               `json:"shopid,omitempty"`
	NormalStock                           int               `json:"normal_stock,omitempty"`
	InstallmentPlans                      []InstallmentPlan `json:"installment_plans,omitempty"`
	ViewCount                             int               `json:"view_count,omitempty"`
	CurrentPromotionHasReserveStock       bool              `json:"current_promotion_has_reserve_stock,omitempty"`
	LikedCount                            int               `json:"liked_count,omitempty"`
	ShowOfficialShopLabel                 bool              `json:"show_official_shop_label,omitempty"`
	PriceMinBeforeDiscount                int               `json:"price_min_before_discount,omitempty"`
	ShowDiscount                          int               `json:"show_discount,omitempty"`
	PreviewInfo                           interface{}       `json:"preview_info,omitempty"`
	Flag                                  int               `json:"flag,omitempty"`
	ExclusivePriceInfo                    interface{}       `json:"exclusive_price_info,omitempty"`
	CurrentPromotionReservedStock         int               `json:"current_promotion_reserved_stock,omitempty"`
	WholesaleTierList                     []interface{}     `json:"wholesale_tier_list,omitempty"`
	GroupBuyInfo                          interface{}       `json:"group_buy_info,omitempty"`
	ShopeeVerified                        bool              `json:"shopee_verified,omitempty"`
	HiddenPriceDisplay                    interface{}       `json:"hidden_price_display,omitempty"`
	TransparentBackgroundImage            string            `json:"transparent_background_image,omitempty"`
	WelcomePackageInfo                    interface{}       `json:"welcome_package_info,omitempty"`
	DiscountStock                         interface{}       `json:"discount_stock,omitempty"`
	CoinInfo                              *CoinInfo         `json:"coin_info,omitempty"`
	IsAdult                               bool              `json:"is_adult,omitempty"`
	Currency                              string            `json:"currency,omitempty"`
	RawDiscount                           int               `json:"raw_discount,omitempty"`
	IsPreferredPlusSeller                 bool              `json:"is_preferred_plus_seller,omitempty"`
	IsCategoryFailed                      bool              `json:"is_category_failed,omitempty"`
	PriceMin                              int               `json:"price_min,omitempty"`
	CanUseBundleDeal                      bool              `json:"can_use_bundle_deal,omitempty"`
	CbOption                              int               `json:"cb_option,omitempty"`
	Brand                                 string            `json:"brand,omitempty"`
	Stock                                 int               `json:"stock,omitempty"`
	Status                                int               `json:"status,omitempty"`
	BundleDealID                          int               `json:"bundle_deal_id,omitempty"`
	IsGroupBuyItem                        bool              `json:"is_group_buy_item,omitempty"`
	Description                           string            `json:"description,omitempty"`
	FlashSale                             *Item             `json:"flash_sale,omitempty"`
	Models                                []Model           `json:"models,omitempty"`
	HasLowFulfillmentRate                 bool              `json:"has_low_fulfillment_rate,omitempty"`
	Price                                 int               `json:"price,omitempty"`
	ShopLocation                          string            `json:"shop_location,omitempty"`
	TierVariations                        []TierVariation   `json:"tier_variations,omitempty"`
	Makeups                               interface{}       `json:"makeups,omitempty"`
	WelcomePackageType                    int               `json:"welcome_package_type,omitempty"`
	ShowOfficialShopLabelInNormalPosition interface{}       `json:"show_official_shop_label_in_normal_position,omitempty"`
	ItemType                              int               `json:"item_type,omitempty"`
	Version                               string            `json:"version,omitempty"`
	Data                                  *Data             `json:"data,omitempty"`
	ErrorMsg                              interface{}       `json:"error_msg,omitempty"`
	Error                                 int               `json:"error,omitempty"`
	BrandSaleBrandCustomLogo              interface{}       `json:"brand_sale_brand_custom_logo,omitempty"`
	Voucher                               interface{}       `json:"voucher,omitempty"`
	FlashSaleType                         int               `json:"flash_sale_type,omitempty"`
	PromoOverlayImage                     string            `json:"promo_overlay_image,omitempty"`
	Modelids                              interface{}       `json:"modelids,omitempty"`
	PromoImages                           []string          `json:"promo_images,omitempty"`
	Promotionid                           int               `json:"promotionid,omitempty"`
	StartTime                             int               `json:"start_time,omitempty"`
	ReminderCount                         interface{}       `json:"reminder_count,omitempty"`
	FlashCatid                            int               `json:"flash_catid,omitempty"`
	IsShopOfficial                        interface{}       `json:"is_shop_official,omitempty"`
	FlashSaleStock                        int               `json:"flash_sale_stock,omitempty"`
	CatLabel                              int               `json:"cat_label,omitempty"`
	EndTime                               int               `json:"end_time,omitempty"`
	IsShopPreferred                       interface{}       `json:"is_shop_preferred,omitempty"`
	PromoName                             string            `json:"promo_name,omitempty"`
	Modelid                               int               `json:"modelid,omitempty"`
	// can be string or int
	ItemGroupID             interface{}    `json:"item_group_id,omitempty"`
	Quantity                int            `json:"quantity,omitempty"`
	AddOnDealInfo           *AddOnDealInfo `json:"add_on_deal_info,omitempty"`
	VideoInfoList           []VideoInfo    `json:"video_info_list,omitempty"`
	IsAddOnSubItem          bool           `json:"is_add_on_sub_item,omitempty"`
	OpcExtraData            *OpcExtraData  `json:"opc_extra_data,omitempty"`
	PromotionID             int            `json:"promotion_id,omitempty"`
	AddOnDealID             int            `json:"add_on_deal_id,omitempty"`
	Offerid                 int            `json:"offerid,omitempty"`
	Source                  string         `json:"source,omitempty"`
	Checkout                bool           `json:"checkout,omitempty"`
	NoneShippableFullReason string         `json:"none_shippable_full_reason,omitempty"`
	IsFlashSale             bool           `json:"is_flash_sale,omitempty"`
	Shippable               bool           `json:"shippable,omitempty"`
	NoneShippableReason     string         `json:"none_shippable_reason,omitempty"`
	ModelName               string         `json:"model_name,omitempty"`
	AppliedPromotionID      int            `json:"applied_promotion_id,omitempty"`
	CartItemChangeTime      int            `json:"cart_item_change_time,omitempty"`

	sh *Shopee
}

type AddOnDealInfo struct {
	AddOnDealID    int    `json:"add_on_deal_id,omitempty"`
	AddOnDealLabel string `json:"add_on_deal_label,omitempty"`
	SubType        int    `json:"sub_type,omitempty"`
}

type VideoInfo struct {
	Duration int    `json:"duration,omitempty"`
	VideoID  string `json:"video_id,omitempty"`
	Version  int    `json:"version,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`
}

type InstallmentPlan struct {
	Banks       []Bank `json:"banks,omitempty"`
	ChannelName string `json:"channel_name,omitempty"`
	IsCc        bool   `json:"is_cc,omitempty"`
	Plans       []Plan `json:"plans,omitempty"`
	ChannelIC   string `json:"channel_ic,omitempty"`
}

type Plan struct {
	Duration       int         `json:"duration,omitempty"`
	DisabledReason interface{} `json:"disabled_reason,omitempty"`
	InterestRate   float64     `json:"interest_rate,omitempty"`
	MonthlyPayment int         `json:"monthly_payment,omitempty"`
	PlanName       string      `json:"plan_name,omitempty"`
}

type Bank struct {
	BankName              string                  `json:"bank_name,omitempty"`
	SubOptions            []Option                `json:"sub_options,omitempty"`
	BankID                int                     `json:"bank_id,omitempty"`
	BankLogo              string                  `json:"bank_logo,omitempty"`
	Description           string                  `json:"description,omitempty"`
	OptionInfo            string                  `json:"option_info,omitempty"`
	IsVirtualAccount      bool                    `json:"is_virtual_account,omitempty"`
	Enabled               bool                    `json:"enabled,omitempty"`
	EligibleTransfer      string                  `json:"eligible_transfer,omitempty"`
	IconBackground        interface{}             `json:"icon_background,omitempty"`
	SubDescriptionInfo    *BankSubDescriptionInfo `json:"sub_description_info,omitempty"`
	PopupConfirmationData *PopupConfirmationData  `json:"popup_confirmation_data,omitempty"`
	OptionIcon            string                  `json:"option_icon,omitempty"`
	Icon                  string                  `json:"icon,omitempty"`
	DisabledReasonKey     string                  `json:"disabled_reason_key,omitempty"`
	DisabledReasonData    *DisabledReasonData     `json:"disabled_reason_data,omitempty"`
	DisabledReason        string                  `json:"disabled_reason,omitempty"`
}

type Option struct {
	DisabledReason string `json:"disabled_reason,omitempty"`
	Data           *Data  `json:"data,omitempty"`
	Name           string `json:"name,omitempty"`
	OptionInfo     string `json:"option_info,omitempty"`
	OptionID       int    `json:"option_id,omitempty"`
	ShortName      string `json:"short_name,omitempty"`
	Description    string `json:"description,omitempty"`
}

type Data struct {
	CartItem                       *Item                    `json:"cart_item,omitempty"`
	ProblematicItems               interface{}              `json:"problematic_items,omitempty"`
	SwitchFulfillmentSourceText    interface{}              `json:"switch_fulfillment_source_text,omitempty"`
	BankName                       string                   `json:"bank_name,omitempty"`
	DownPayment                    int                      `json:"down_payment,omitempty"`
	Name                           string                   `json:"name,omitempty"`
	InterestRate                   float64                  `json:"interest_rate,omitempty"`
	OptionID                       interface{}              `json:"option_id,omitempty"`
	BankID                         int                      `json:"bank_id,omitempty"`
	InstallmentAmount              int                      `json:"installment_amount,omitempty"`
	ChannelID                      interface{}              `json:"channel_id,omitempty"`
	MonthlyInstallment             int                      `json:"monthly_installment,omitempty"`
	Tenure                         int                      `json:"tenure,omitempty"`
	TotalAmount                    int                      `json:"total_amount,omitempty"`
	RebateRoundingFactor           int                      `json:"rebate_rounding_factor,omitempty"`
	BuyerLocationGroupID           int                      `json:"buyer_location_group_id,omitempty"`
	BuyerListEntityPermissionGroup int                      `json:"buyer_list_entity_permission_group,omitempty"`
	SellerLocationGroupID          int                      `json:"seller_location_group_id,omitempty"`
	PercentageRebateCap            int                      `json:"percentage_rebate_cap,omitempty"`
	PercentageDiscountCap          int                      `json:"percentage_discount_cap,omitempty"`
	MinOrderTotal                  int                      `json:"min_order_total,omitempty"`
	EntityPermissionGroup          int                      `json:"entity_permission_group,omitempty"`
	DiscountRoundingFactor         int                      `json:"discount_rounding_factor,omitempty"`
	PaymentType                    int                      `json:"payment_type,omitempty"`
	CustomLabel1                   string                   `json:"custom_label1,omitempty"`
	CustomLabel2                   string                   `json:"custom_label2,omitempty"`
	AllPromotionRules              []AllPromotionRule       `json:"all_promotion_rules,omitempty"`
	ShopOrders                     []ShopOrder              `json:"shop_orders,omitempty"`
	Fsv                            *Fsv                     `json:"fsv,omitempty"`
	ShopOrderIDS                   []ShopOrder              `json:"shop_order_ids,omitempty"`
	LogisticsChannels              []LogisticsChannel       `json:"logistics_channels,omitempty"`
	AddOnDeals                     []interface{}            `json:"add_on_deals,omitempty"`
	IsFreeShippingVoucherToggledOn bool                     `json:"is_free_shipping_voucher_toggled_on,omitempty"`
	CoinMinRedeem                  int                      `json:"coin_min_redeem,omitempty"`
	BuyerLocationGroupIDS          []int                    `json:"buyer_location_group_ids,omitempty"`
	ShopOrderIDList                []ShopOrder              `json:"shop_order_id_list,omitempty"`
	DiscountTabShopOrderIDS        *DiscountTabShopOrderIDS `json:"discount_tab_shop_order_ids,omitempty"`
	DisableTab                     bool                     `json:"disable_tab,omitempty"`
	ShowTabs                       []string                 `json:"show_tabs,omitempty"`
	Shoporders                     []interface{}            `json:"shoporders,omitempty"`
	CanUseCoins                    bool                     `json:"can_use_coins,omitempty"`
	UseCoins                       bool                     `json:"use_coins,omitempty"`
	CoinDiscount                   int                      `json:"coin_discount,omitempty"`
	CoinInfo                       *CoinInfo                `json:"coin_info,omitempty"`
	CoinText                       string                   `json:"coin_text,omitempty"`
	CoinDetailedText               string                   `json:"coin_detailed_text,omitempty"`
	Taxes                          *Taxes                   `json:"taxes,omitempty"`
	CardPromotionID                int                      `json:"card_promotion_id,omitempty"`
	CardPromotionDiscount          int                      `json:"card_promotion_discount,omitempty"`
	CardPromotionEnabled           bool                     `json:"card_promotion_enabled,omitempty"`
	Description                    string                   `json:"description,omitempty"`
	PromotionMsg                   string                   `json:"promotion_msg,omitempty"`
	InvalidMessageCode             int                      `json:"invalid_message_code,omitempty"`
	InvalidMessage                 string                   `json:"invalid_message,omitempty"`
	PlatformVouchers               []interface{}            `json:"platform_vouchers,omitempty"`
	ShopVouchers                   []interface{}            `json:"shop_vouchers,omitempty"`
	FreeShippingVoucherInfo        *FreeShippingVoucherInfo `json:"free_shipping_voucher_info,omitempty"`
	TotalPayment                   []int                    `json:"total_payment,omitempty"`
	BundleDeals                    []interface{}            `json:"bundle_deals,omitempty"`
	BundleDealsPromotionDiscount   int                      `json:"bundle_deals_promotion_discount,omitempty"`
	ShopVoucherEntrance            []interface{}            `json:"shop_voucher_entrance,omitempty"`
	Error                          int                      `json:"error,omitempty"`
	ErrorMessage                   string                   `json:"error_message,omitempty"`
	MessageLevel                   *MessageLevel            `json:"message_level,omitempty"`
	WarnMessage                    interface{}              `json:"warn_message,omitempty"`
	MinAmount                      int                      `json:"min_amount,omitempty"`
}

type Taxes struct {
	Orders []int `json:"orders,omitempty"`
	Total  int   `json:"total,omitempty"`
}

type FreeShippingVoucherInfo struct {
	FreeShippingVoucherID   int    `json:"free_shipping_voucher_id,omitempty"`
	DisabledReason          string `json:"disabled_reason,omitempty"`
	Description             string `json:"description,omitempty"`
	FreeShippingVoucherCode string `json:"free_shipping_voucher_code,omitempty"`
}

type AllPromotionRule struct {
	DiscountDelta    int      `json:"discount_delta,omitempty"`
	ChannelIDS       []string `json:"channel_ids,omitempty"`
	Shopid           int      `json:"shopid,omitempty"`
	DiscountRuleFlag int      `json:"discount_rule_flag,omitempty"`
	RebateFlag       int      `json:"rebate_flag,omitempty"`
	Priority         int      `json:"priority,omitempty"`
	ExtraData        *Data    `json:"extra_data,omitempty"`
	ID               int      `json:"id,omitempty"`
	DiscountFlag     int      `json:"discount_flag,omitempty"`
}

type ShopOrder struct {
	ShopHasVoucher                      interface{}         `json:"shop_has_voucher,omitempty"`
	Shop                                *Shop               `json:"shop,omitempty"`
	Items                               []Item              `json:"items,omitempty"`
	Shopid                              int                 `json:"shopid,omitempty"`
	AddinTime                           int                 `json:"addin_time,omitempty"`
	ClickTime                           interface{}         `json:"click_time,omitempty"`
	ItemBriefs                          []ItemBrief         `json:"item_briefs,omitempty"`
	ShopVouchers                        []interface{}       `json:"shop_vouchers,omitempty"`
	Logistics                           *Logistics          `json:"logistics,omitempty"`
	BuyerAddressData                    *PaymentChannelData `json:"buyer_address_data,omitempty"`
	SelectedPreferredDeliveryTimeSlotID interface{}         `json:"selected_preferred_delivery_time_slot_id,omitempty"`
}

type Logistics struct {
	RecommendedChannelids           interface{}                `json:"recommended_channelids,omitempty"`
	NonIntegratedChannelids         []interface{}              `json:"non_integrated_channelids,omitempty"`
	LogisticServiceTypes            *LogisticServiceTypes      `json:"logistic_service_types,omitempty"`
	LogisticChannels                map[string]LogisticChannel `json:"logistic_channels,omitempty"`
	IntegratedChannelids            []int                      `json:"integrated_channelids,omitempty"`
	VoucherWalletCheckingChannelIDS []interface{}              `json:"voucher_wallet_checking_channel_ids,omitempty"`
}

type PaymentChannelData struct {
}

type Shop struct {
	Shopname      string `json:"shopname,omitempty"`
	HolidayModeOn bool   `json:"holiday_mode_on,omitempty"`
	Shopid        int    `json:"shopid,omitempty"`
	Username      string `json:"username,omitempty"`
	Status        int    `json:"status,omitempty"`
	// can be int or bool
	CbOption              interface{} `json:"cb_option,omitempty"`
	ShowOfficialShopLabel bool        `json:"show_official_shop_label,omitempty"`
	Portrait              string      `json:"portrait,omitempty"`
	FollowingCount        int         `json:"following_count,omitempty"`
	Userid                int         `json:"userid,omitempty"`
	IsFreeShipping        bool        `json:"is_free_shipping,omitempty"`
	IsShopeeVerified      bool        `json:"is_shopee_verified,omitempty"`
	HasWelcomePackageItem bool        `json:"has_welcome_package_item,omitempty"`
	EnabledChannelids     []string    `json:"enabled_channelids,omitempty"`
	PromotionRules        []int       `json:"promotion_rules,omitempty"`
	ShopTag               int         `json:"shop_tag,omitempty"`
	AddinTime             int         `json:"addin_time,omitempty"`
	ClickTime             int         `json:"click_time,omitempty"`
	RemarkType            int         `json:"remark_type,omitempty"`
	SupportEreceipt       bool        `json:"support_ereceipt,omitempty"`
	Images                string      `json:"images,omitempty"`
	IsOfficialShop        bool        `json:"is_official_shop,omitempty"`
	ShopName              string      `json:"shop_name,omitempty"`
}

type Fsv struct {
	FsvMessage string      `json:"fsv_message,omitempty"`
	URL        string      `json:"url,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	ErrorMsg   interface{} `json:"error_msg,omitempty"`
}

type ItemBrief struct {
	Itemid  int `json:"itemid,omitempty"`
	Modelid int `json:"modelid,omitempty"`
	// can be string or int
	ItemGroupID        interface{} `json:"item_group_id,omitempty"`
	Quantity           int         `json:"quantity,omitempty"`
	AppliedPromotionID int         `json:"applied_promotion_id,omitempty"`
	OfferID            int         `json:"offerid,omitempty"`
	Price              int         `json:"price,omitempty"`
	IsAddOnSubItem     bool        `json:"is_add_on_sub_item,omitempty"`
	AddOnDealID        int         `json:"add_on_deal_id,omitempty"`
	Status             int         `json:"status,omitempty"`
	CartItemChangeTime int         `json:"cart_item_change_time,omitempty"`
}

type LogisticsChannel struct {
	Category    int    `json:"category,omitempty"`
	MinimumCost int    `json:"minimum_cost,omitempty"`
	Name        string `json:"name,omitempty"`
	ChannelID   int    `json:"channel_id,omitempty"`
}

type DiscountTabShopOrderIDS struct {
	Discount []interface{} `json:"discount,omitempty"`
	Other    []interface{} `json:"other,omitempty"`
}

type Attribute struct {
	IsPendingQc bool   `json:"is_pending_qc,omitempty"`
	Idx         int    `json:"idx,omitempty"`
	Value       string `json:"value,omitempty"`
	ID          int    `json:"id,omitempty"`
	IsTimestamp bool   `json:"is_timestamp,omitempty"`
	Name        string `json:"name,omitempty"`
}

type Category struct {
	DisplayName        string      `json:"display_name,omitempty"`
	Catid              int         `json:"catid,omitempty"`
	Image              interface{} `json:"image,omitempty"`
	NoSub              bool        `json:"no_sub,omitempty"`
	IsDefaultSubcat    bool        `json:"is_default_subcat,omitempty"`
	BlockBuyerPlatform interface{} `json:"block_buyer_platform,omitempty"`
	Catids             []int       `json:"catids,omitempty"`
}

type CoinInfo struct {
	SpendCashUnit     int           `json:"spend_cash_unit,omitempty"`
	CoinEarnItems     []interface{} `json:"coin_earn_items,omitempty"`
	CoinOffset        int           `json:"coin_offset,omitempty"`
	CoinEarn          int           `json:"coin_earn,omitempty"`
	CoinEarnByVoucher int           `json:"coin_earn_by_voucher,omitempty"`
	CoinUsed          int           `json:"coin_used,omitempty"`
}

type Rating struct {
	RatingStar        float64 `json:"rating_star,omitempty"`
	RatingCount       []int   `json:"rating_count,omitempty"`
	RcountWithImage   int     `json:"rcount_with_image,omitempty"`
	RcountWithContext int     `json:"rcount_with_context,omitempty"`
}

type Model struct {
	Itemid                          int           `json:"itemid,omitempty"`
	Status                          int           `json:"status,omitempty"`
	CurrentPromotionReservedStock   int           `json:"current_promotion_reserved_stock,omitempty"`
	Name                            string        `json:"name,omitempty"`
	Promotionid                     int           `json:"promotionid,omitempty"`
	Price                           int           `json:"price,omitempty"`
	PriceStocks                     []PriceStock  `json:"price_stocks,omitempty"`
	CurrentPromotionHasReserveStock bool          `json:"current_promotion_has_reserve_stock,omitempty"`
	Currency                        string        `json:"currency,omitempty"`
	NormalStock                     int           `json:"normal_stock,omitempty"`
	Extinfo                         *ModelExtinfo `json:"extinfo,omitempty"`
	PriceBeforeDiscount             int           `json:"price_before_discount,omitempty"`
	Modelid                         int           `json:"modelid,omitempty"`
	Sold                            int           `json:"sold,omitempty"`
	Stock                           int           `json:"stock,omitempty"`
}

type ModelExtinfo struct {
	SellerPromotionLimit       int         `json:"seller_promotion_limit,omitempty"`
	HasShopeePromo             bool        `json:"has_shopee_promo,omitempty"`
	GroupBuyInfo               interface{} `json:"group_buy_info,omitempty"`
	HolidayModeOldStock        interface{} `json:"holiday_mode_old_stock,omitempty"`
	TierIndex                  []int       `json:"tier_index,omitempty"`
	SellerPromotionRefreshTime int         `json:"seller_promotion_refresh_time,omitempty"`
}

type PriceStock struct {
	ModelID                  int           `json:"model_id,omitempty"`
	StockoutTime             int           `json:"stockout_time,omitempty"`
	Region                   string        `json:"region,omitempty"`
	Rebate                   interface{}   `json:"rebate,omitempty"`
	Price                    int           `json:"price,omitempty"`
	PromotionType            int           `json:"promotion_type,omitempty"`
	AllocatedStock           int           `json:"allocated_stock,omitempty"`
	ShopID                   int           `json:"shop_id,omitempty"`
	EndTime                  int           `json:"end_time,omitempty"`
	StockBreakdownByLocation []interface{} `json:"stock_breakdown_by_location,omitempty"`
	ItemID                   int           `json:"item_id,omitempty"`
	PromotionID              int           `json:"promotion_id,omitempty"`
	PurchaseLimit            int           `json:"purchase_limit,omitempty"`
	StartTime                int           `json:"start_time,omitempty"`
	Stock                    int           `json:"stock,omitempty"`
}

type TierVariation struct {
	Images     []string      `json:"images,omitempty"`
	Properties []interface{} `json:"properties,omitempty"`
	Type       int           `json:"type,omitempty"`
	Name       string        `json:"name,omitempty"`
	Options    []string      `json:"options,omitempty"`
}

// ItemByIDAndShopID get item by item id and shop id
func (sh *Shopee) ItemByIDAndShopID(id, shopID int) (*Item, error) {
	raw, err := sh.get("/v2/item/get", url.Values{
		"itemid": {strconv.Itoa(id)},
		"shopid": {strconv.Itoa(shopID)},
	})

	if err != nil {
		return nil, err
	}

	item := new(Item)
	if err := json.Unmarshal(raw, item); err != nil {
		return nil, err
	}

	item.Item.sh = sh

	return item.Item, nil
}

// ItemByLink get item by link.
// Example link:
//  https://shopee.co.id/Asus-ROG-Phone-3-ZS661KS-6A007ID-128GB-8GB-Black-i.155149633.6554463078
func (sh *Shopee) ItemByLink(link string) (*Item, error) {
	// check if link is in valid form or not
	uri, err := url.Parse(link)
	if err != nil {
		return nil, errors.New("invalid item link")
	}

	if !strings.Contains(uri.Host, "shopee.") {
		return nil, errors.New("invalid item link")
	}

	// extract item id and shop id
	split := strings.Split(link, ".")
	if len(split)-2 <= 0 {
		return nil, errors.New("invalid item link")
	}

	shopIDStr := split[len(split)-2]
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		return nil, errors.New("invalid item link")
	}

	itemIDStr := split[len(split)-1]
	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		return nil, errors.New("invalid item link")
	}

	return sh.ItemByIDAndShopID(itemID, shopID)
}

// SetShopee set client to item
func (i *Item) SetShopee(sh *Shopee) {
	i.sh = sh
}

// AddToCart add item to cart.
// Increase qty for increase quantities.
// Item can have different variations or models, so make sure to define which modelID you want to add.
// If modelID == 0, the first model in Item.Models will be used.
// Will use Add On Deal if present.
func (i *Item) AddToCart(modelID, qty int) (*Item, error) {
	body := map[string]interface{}{
		"checkout":             true,
		"client_source":        1,
		"donot_add_quantity":   false,
		"itemid":               i.Itemid,
		"modelid":              modelID,
		"quantity":             qty,
		"shopid":               i.Shopid,
		"source":               `{"refer_urls":[]}`,
		"update_checkout_only": false,
	}

	if modelID == 0 {
		if len(i.Models) > 0 {
			body["modelid"] = i.Models[0].Modelid
		}

		if i.Modelid != 0 {
			body["modelid"] = i.Modelid
		}
	}

	if i.AddOnDealInfo != nil {
		body["add_on_deal_id"] = i.AddOnDealInfo.AddOnDealID
	}

	refer := composeItemURL(i)
	raw, err := i.sh.post("/v2/cart/add_to_cart", refer, body)
	if err != nil {
		return nil, err
	}

	item := new(Item)
	if err := json.Unmarshal(raw, item); err != nil {
		return nil, err
	}

	if item.Error > 0 {
		return nil, errors.New(string(raw))
	}

	return item, nil
}

func composeItemURL(item *Item) string {
	var urlStr string
	for _, c := range strings.Split(item.Name, "") {
		if !isCharAChar(c) && !isCharANum(c) {
			urlStr += "-"
			continue
		}

		urlStr += c
	}

	urlStr += "-i"
	urlStr = "https://shopee.co.id/" + urlStr + "." + strconv.Itoa(item.Shopid) + "." + strconv.Itoa(item.Itemid)

	return urlStr
}

func isCharAChar(c string) bool {
	latin := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	for _, char := range latin {
		if c == char {
			return true
		}
	}

	return false
}

func isCharANum(c string) bool {
	num := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, char := range num {
		if c == char {
			return true
		}
	}

	return false
}
