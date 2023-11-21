package gotiktok

type SearchReturnRequest struct {
	Locale   *string  `json:"locale" validate:"omitempty,oneof=en-US id-ID"`
	OrderIds []string `json:"order_ids" validate:"omitempty"`
}

type ReturnOrder struct {
	ArbitrationStatus        string                     `json:"arbitration_status"`
	CreateTime               int                        `json:"create_time"`
	HandoverMethod           string                     `json:"handover_method"`
	NextReturnId             string                     `json:"next_return_id"`
	OrderId                  string                     `json:"order_id"`
	PreReturnId              string                     `json:"pre_return_id"`
	RefundAmount             RefundAmount               `json:"refund_amount"`
	ReturnID                 string                     `json:"return_id"`
	ReturnLineItems          []ReturnLineItems          `json:"return_line_items"`
	ReturnProviderId         string                     `json:"return_provider_id"`
	ReturnProviderName       string                     `json:"return_provider_name"`
	ReturnReason             string                     `json:"return_reason"`
	ReturnReasonText         string                     `json:"return_reason_text"`
	ReturnStatus             string                     `json:"return_status"`
	ReturnTrackingNumber     string                     `json:"return_tracking_number"`
	ReturnType               string                     `json:"return_type"`
	Role                     string                     `json:"role"`
	SellerNextActionResponse []SellerNextActionResponse `json:"seller_next_action_response"`
	ShipmentType             string                     `json:"shipment_type"`
	UpdateTime               int                        `json:"update_time"`
}

type RefundAmount struct {
	Currency          string `json:"currency"`
	RefundShippingFee string `json:"refund_shipping_fee"`
	RefundSubtotal    string `json:"refund_subtotal"`
	RefundTax         string `json:"refund_tax"`
	RefundTotal       string `json:"refund_total"`
}

type ProductImage struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type ReturnLineItems struct {
	OrderLineItemId  string       `json:"order_line_item_id"`
	ProductImage     ProductImage `json:"product_image"`
	ProductName      string       `json:"product_name"`
	RefundAmount     RefundAmount `json:"refund_amount"`
	ReturnLineItemId string       `json:"return_line_item_id"`
	SellerSku        string       `json:"seller_sku"`
	SkuId            string       `json:"sku_id"`
	SkuName          string       `json:"sku_name"`
}

type SellerNextActionResponse struct {
	Action   string `json:"action"`
	Deadline int    `json:"deadline"`
}

type SearchReturnResponse struct {
	ReturnOrders  []ReturnOrder `json:"return_orders"`
	TotalCount    int           `json:"total_count"`
	NextPageToken string        `json:"next_page_token"`
}
