package gotiktok

type Shop struct {
	Id         string `json:"id"`
	Code       string `json:"code"`
	Cipher     string `json:"cipher"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	SellerType string `json:"seller_type"`
}

type ShopList struct {
	Shops []Shop `json:"shops"`
}
