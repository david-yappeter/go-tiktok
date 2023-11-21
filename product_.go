package gotiktok

type Category struct {
	Id        string `json:"id"`
	ParentId  string `json:"parent_id"`
	LocalName string `json:"local_name"`
	IsLeaf    bool   `json:"is_leaf"`
}

type Brand struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Inventory struct {
	WarehouseId string `json:"warehouse_id,omitempty"`
	Quantity    int    `json:"quantity"`
}

type Image struct {
	Height    int      `json:"height"`
	Width     int      `json:"width"`
	ThumbUrls []string `json:"thumb_urls"`
	Uri       string   `json:"uri"`
	Urls      []string `json:"urls"`
}

type Value struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProductAttribute struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

type Price struct {
	Currency          string `json:"currency"`
	SalePrice         string `json:"sale_price"`
	TaxExclusivePrice string `json:"tax_exclusive_price"`
}

type SalesAttr struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ValueId   string `json:"value_id"`
	ValueName string `json:"value_name"`
	SkuImg    Image  `json:"sku_img"`
}

type SKUData struct {
	Id              string      `json:"id"`
	SellerSku       string      `json:"seller_sku"`
	Price           Price       `json:"price"`
	Inventory       []Inventory `json:"inventory"`
	SalesAttributes []SalesAttr `json:"sales_attributes"`
}

type PackageDimensions struct {
	Height string `json:"height"`
	Length string `json:"length"`
	Unit   string `json:"unit"`
	Width  string `json:"width"`
}

type PackageWeight struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

type ProductDetailData struct {
	Id                string             `json:"id"`
	Status            string             `json:"status"`
	Title             string             `json:"title"`
	CategoryChains    []Category         `json:"category_chains"`
	Brand             Brand              `json:"brand"`
	MainImages        []Image            `json:"main_images"`
	PackageDimensions PackageDimensions  `json:"package_dimensions"`
	PackageWeight     PackageWeight      `json:"package_weight"`
	IsCodAllowed      bool               `json:"is_cod_allowed"`
	Description       string             `json:"description"`
	UpdateTime        int64              `json:"update_time"`
	CreateTime        int64              `json:"create_time"`
	ProductAttributes []ProductAttribute `json:"product_attributes"`
	Skus              []SKUData          `json:"skus"`
}

type SearchProductRequest struct {
	Status *string `json:"status" validate:"omitempty,oneof=ALL DRAFT PENDING FAILED ACTIVATE SELLER_DEACTIVATED PLATFORM_DEACTIVATED FREEZE DELETED"`
}

type ProductData struct {
	Id           string    `json:"id"`
	Status       string    `json:"status"`
	Title        string    `json:"title"`
	Skus         []SKUData `json:"skus"`
	SalesRegions []string  `json:"sales_regions"`
	CreateTime   int64     `json:"create_time"`
	UpdateTime   int64     `json:"update_time"`
}

type ProductList struct {
	Products      []ProductData `json:"products"`
	TotalCount    int           `json:"total_count"`
	NextPageToken string        `json:"next_page_token"`
}
