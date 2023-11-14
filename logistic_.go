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
