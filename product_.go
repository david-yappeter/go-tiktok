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

type CreateProductRequest struct {
	SaveMode *string `json:"save_mode"`
	/*
		Valid Values:
		AS_DRAFT, LISTING
		DEFAULT: LISTING
	*/
	Description       string                                 `json:"description"`
	CategoryId        string                                 `json:"category_id"`
	BrandId           *string                                `json:"brand_id"`
	MainImages        []CreateProductRequestMainImage        `json:"main_images"`
	Skus              []CreateProductRequestSku              `json:"skus"`
	Title             string                                 `json:"title"`
	IsCodAllowed      bool                                   `json:"is_cod_allowed"`
	PackageDimensions *PackageDimensions                     `json:"package_dimensions"`
	ProductAttributes []CreateProductRequestProductAttribute `json:"product_attribute"`
	PackageWeight     PackageWeight                          `json:"package_weight"`
	Video             *CreateProductRequestVideo             `json:"video"`
	SizeChart         *CreateProductRequestSizeChart         `json:"size_chart"`
}

type CreateProductRequestMainImage struct {
	Uri string `json:"uri"`
}

type CreateProductRequestSku struct {
	Inventory []CreateProductRequestSkuInventory `json:"inventory"`
	SellerSku *string                            `json:"seller_sku"`

	Price CreateProductRequestSkuPrice `json:"price"`
}

type CreateProductRequestSkuInventory struct {
	WarehouseId string `json:"warehouse_id"`
	Quantity    int    `json:"quantity"`
}

type CreateProductRequestSkuPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type CreateProductRequestSizeChart struct {
	Image *CreateProductRequestSizeChartImage `json:"image"`
}

type CreateProductRequestSizeChartImage struct {
	Uri string `json:"uri"`
}

type CreateProductRequestProductAttribute struct {
	Id     string                                      `json:"id"`
	Values []CreateProductRequestProductAttributeValue `json:"values"`
}

type CreateProductRequestProductAttributeValue struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

type CreateProductRequestPackageWeight struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type CreateProductRequestVideo struct {
	Id string `json:"id"`
}

type CreateProductRequestProductVideo struct {
	VideoId string `json:"video_id"`
}

type UpdateProductStockRequest struct {
	ProductId string `json:"product_id"`
}

type UpdateProductStockRequestSku struct {
	Id         string                                  `json:"id"`
	StockInfos []UpdateProductStockRequestSkuStockInfo `json:"stock_infos"`
}

type UpdateProductStockRequestSkuStockInfo struct {
	WarehouseId    string `json:"warehouse_id"`
	AvailableStock int    `json:"available_stock"`
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

type CreateProductResponse struct {
	ProductId string                         `json:"product_id"`
	Skus      []CreateProductResponseSku     `json:"skus"`
	Warnings  []CreateProductResponseWarning `json:"warnings"`
}

type CreateProductResponseSku struct {
	Id              string                                   `json:"id"`
	SellerSku       string                                   `json:"seller_sku"`
	SalesAttributes []CreateProductResponseSkuSalesAttribute `json:"sales_attributes"`
	ExternalSkuId   string                                   `json:"external_sku_id"`
}

type CreateProductResponseSkuSalesAttribute struct {
	AttributeId string `json:"attribute_id"`
	ValueId     string `json:"value_id"`
}

type CreateProductResponseWarning struct {
	Message string `json:"message"`
}

type UpdateProductInventoryRequest struct {
	Skus []UpdateProductInventoryRequestSku `json:"skus"`
}

type UpdateProductInventoryRequestSku struct {
	Id        string                                      `json:"id"`
	Inventory []UpdateProductInventoryRequestSkuInventory `json:"inventory"`
}

type UpdateProductInventoryRequestSkuInventory struct {
	WarehouseId string `json:"warehouse_id"`
	Quantity    int    `json:"quantity"`
}

type UpdateProductInventoryResponse struct {
	Errors []UpdateProductInventoryResponseError `json:"errors"`
}

type UpdateProductInventoryResponseError struct {
	Code    int                                       `json:"code"`
	Message string                                    `json:"message"`
	Detail  UpdateProductInventoryResponseErrorDetail `json:"detail"`
}

type UpdateProductInventoryResponseErrorDetail struct {
	SkuId       string                                                `json:"sku_id"`
	ExtraErrors []UpdateProductInventoryResponseErrorDetailExtraError `json:"extra_errors"`
}

type UpdateProductInventoryResponseErrorDetailExtraError struct {
	WarehouseId string `json:"warehouse_id"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
}

type UpdateProductPriceRequest struct {
	Skus []UpdateProductPriceRequestSku `json:"skus"`
}

type UpdateProductPriceRequestSku struct {
	Id    string                            `json:"id"`
	Price UpdateProductPriceRequestSkuPrice `json:"price"`
}

type UpdateProductPriceRequestSkuPrice struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type UpdateProductPriceResponse struct {
}

type ActivateProductRequest struct {
	ProductIds []string `json:"product_ids"`
}

type ActivateProductResponse struct {
	Errors []ActivateProductResponseError `json:"errors"`
}

type ActivateProductResponseError struct {
	Code    int                                `json:"code"`
	Message string                             `json:"message"`
	Detail  ActivateProductResponseErrorDetail `json:"detail"`
}

type ActivateProductResponseErrorDetail struct {
	ProductId   string                                         `json:"product_id"`
	ExtraErrors []ActivateProductResponseErrorDetailExtraError `json:"extra_errors"`
}

type ActivateProductResponseErrorDetailExtraError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DeactivateProductRequest struct {
	ProductIds []string `json:"product_ids"`
}

type DeactivateProductResponse struct {
	Errors []DeactivateProductResponseError `json:"errors"`
}

type DeactivateProductResponseError struct {
	Code    int                                  `json:"code"`
	Message string                               `json:"message"`
	Detail  DeactivateProductResponseErrorDetail `json:"detail"`
}

type DeactivateProductResponseErrorDetail struct {
	ProductId string `json:"product_id"`
}
