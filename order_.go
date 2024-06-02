package gotiktok

type DistrictInfo struct {
	AddressLevelName string `json:"address_level_name"`
	AddressName      string `json:"address_name"`
}

type RecipientAddress struct {
	AddressDetail string         `json:"address_detail"`
	AddressLine1  string         `json:"address_line_1"`
	AddressLine2  string         `json:"address_line_2"`
	AddressLine3  string         `json:"address_line_3"`
	AddressLine4  string         `json:"address_line_4"`
	FullAddress   string         `json:"full_address"`
	Name          string         `json:"name"`
	PhoneNumber   string         `json:"phone_number"`
	PostalCode    string         `json:"postal_code"`
	RegionCode    string         `json:"region_code"`
	DistrictInfo  []DistrictInfo `json:"district_info"`
}

type Payment struct {
	BuyerServiceFee             string `json:"buyer_service_fee"`
	Currency                    string `json:"currency"`
	OriginalTotalProductPrice   string `json:"original_total_product_price"`
	OriginalShippingFee         string `json:"original_shipping_fee"`
	PlatformDiscount            string `json:"platform_discount"`
	SellerDiscount              string `json:"seller_discount"`
	ShippingFee                 string `json:"shipping_fee"`
	ShippingFeePlatformDiscount string `json:"shipping_fee_platform_discount"`
	ShippingFeeSellerDiscount   string `json:"shipping_fee_seller_discount"`
	SubTotal                    string `json:"sub_total"`
	Tax                         string `json:"tax"`
	TotalAmount                 string `json:"total_amount"`
}

type OrderLineItem struct {
	Id                   string `json:"id"`
	SkuId                string `json:"sku_id"`
	ProductId            string `json:"product_id"`
	ProductName          string `json:"product_name"`
	SkuImage             string `json:"sku_image"`
	PackageId            string `json:"package_id"`
	PackageStatus        string `json:"package_status"`
	IsGift               bool   `json:"is_gift"`
	OriginalPrice        string `json:"original_price"`
	PlatformDiscount     string `json:"platform_discount"`
	RTSTime              int    `json:"rts_time"`
	SalePrice            string `json:"sale_price"`
	SellerSku            string `json:"seller_sku"`
	SellerDiscount       string `json:"seller_discount"`
	ShippingProviderId   string `json:"shipping_provider_id"`
	ShippingProviderName string `json:"shipping_provider_name"`
	SkuType              string `json:"sku_type"`
	TrackingNumber       string `json:"tracking_number"`
}

type OrderPackage struct {
	Id string `json:"id"`
}

type OrderDetail struct {
	Id                          string           `json:"id"`
	BuyerEmail                  string           `json:"buyer_email"`
	BuyerMessage                string           `json:"buyer_message"`
	CancelOrderSLATime          int              `json:"cancel_order_sla_time"`
	CreateTime                  int              `json:"create_time"`
	DeliveryOptionId            string           `json:"delivery_option_id"`
	DeliveryOptionName          string           `json:"delivery_option_name"`
	DeliverySLATime             int              `json:"delivery_sla_time"`
	FulfillmentType             string           `json:"fulfillment_type"`
	HasUpdatedReceipientAddress bool             `json:"has_updated_receipient_address"`
	IsCod                       bool             `json:"is_cod"`
	IsSampleOrder               bool             `json:"is_sample_order"`
	Packages                    []OrderPackage   `json:"packages"`
	LineItems                   []OrderLineItem  `json:"line_items"`
	PaidTime                    int64            `json:"paid_time"`
	Payment                     Payment          `json:"payment,omitempty"`
	RecipientAddress            RecipientAddress `json:"recipient_address,omitempty"`
	RtsSLATime                  int              `json:"rts_sla_time"`
	RtsTime                     int              `json:"rts_time"`
	ShippingProvider            string           `json:"shipping_provider"`
	ShippingProviderID          string           `json:"shipping_provider_id"`
	ShippingType                string           `json:"shipping_type"`
	Status                      string           `json:"status"`
	TrackingNumber              string           `json:"tracking_number"`
	TtsSLATime                  int              `json:"tts_sla_time"`
	UpdateTime                  int              `json:"update_time"`
	UserId                      string           `json:"user_id"`
	WarehouseId                 string           `json:"warehouse_id"`
}

type OrderDetailList struct {
	Orders []OrderDetail `json:"orders"`
}
