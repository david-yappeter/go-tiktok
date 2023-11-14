package gotiktok

type WarehouseAddress struct {
	Region        string `json:"region"`
	State         string `json:"state"`
	City          string `json:"city"`
	District      string `json:"district"`
	Town          string `json:"town"`
	ContactPerson string `json:"contact_person"`
	PostalCode    string `json:"postal_code"`
	FullAddress   string `json:"full_address"`
	RegionCode    string `json:"region_code"`
	PhoneNumber   string `json:"phone_number"`
}

type Warehouse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	EffectStatus string `json:"effect_status"`
	Type         string `json:"type"`
	SubType      string `json:"sub_type"`
	IsDefault    bool   `json:"is_default"`
}

type WarehouseList struct {
	Warehouses []Warehouse `json:"warehouses"`
}

type WarehouseDeliveryOptionDimensionLimit struct {
	MaxHeight int    `json:"max_height"`
	MaxLength int    `json:"max_length"`
	MaxWidth  int    `json:"max_width"`
	Unit      string `json:"unit"`
}

type WarehouseDeliveryOptionWeightLimit struct {
	MaxWeight int    `json:"max_weight"`
	MinWeight int    `json:"min_weight"`
	Unit      string `json:"unit"`
}

type WarehouseDeliveryOption struct {
	Id          string                                `json:"id"`
	Name        string                                `json:"name"`
	Type        string                                `json:"type"`
	Description string                                `json:"description"`
	WeightLimit WarehouseDeliveryOptionWeightLimit    `json:"weight_limit"`
	Dimension   WarehouseDeliveryOptionDimensionLimit `json:"dimension_limit"`
}

type WarehouseDeliveryOptionList struct {
	DeliveryOptions []WarehouseDeliveryOption `json:"delivery_options"`
}

type ShippingProvider struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ShippingProviderList struct {
	ShippingProviders []ShippingProvider `json:"shipping_providers"`
}
