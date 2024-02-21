package gotiktok

type GetCategoriesRequest struct {
	Locale *string `json:"locale"`
	/*
		Currently Supported:
		en-GB, en-US, id-ID, ms-MY, th-TH, vi-VN, zh-CN
	*/
}

type GetCategoryResponse struct {
	Categories []GetCategoryResponseCategory `json:"categories"`
}

type GetCategoryResponseCategory struct {
	Id                 string   `json:"id"`
	ParentId           string   `json:"parent_id"`
	LocalName          string   `json:"local_name"`
	IsLeaf             bool     `json:"is_leaf"`
	PermissionStatuses []string `json:"permission_statuses"`
}

type GetCategoryRulesResponse struct {
	ProductCertifications []GetCategoryRulesResponseProductCertification `json:"product_certifications"`
	SizeChart             GetCategoryRulesResponseSizeChart              `json:"size_chart"`
	Cod                   GetCategoryRulesResponseCod                    `json:"cod"`
	PackageDimension      GetCategoryRulesResponsePackageDimension       `json:"package_dimension"`
}

type GetCategoryRulesResponseProductCertification struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	IsRequired     bool   `json:"is_required"`
	SampleImageUrl string `json:"sample_image_url"`
}

type GetCategoryRulesResponseSizeChart struct {
	IsSupported bool `json:"is_supported"`
	IsRequired  bool `json:"is_required"`
}

type GetCategoryRulesResponseCod struct {
	IsSupported bool `json:"is_supported"`
}

type GetCategoryRulesResponsePackageDimension struct {
	IsRequired bool `json:"is_required"`
}

type GetAttributesRequest struct {
	Locale *string `json:"locale"`
	/*
		Currently Supported:
		en-GB, en-US, id-ID, ms-MY, th-TH, vi-VN, zh-CN
	*/
}

type GetAttributesResponse struct {
	Attributes []GetAttributesResponseAttribute `json:"attributes"`
}

type GetAttributesResponseAttribute struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	/*
		Values:
		SALES_PROPERTY
		PRODUCT_PROPERTY
	*/
	IsRequired          bool                                  `json:"is_requried"`
	Values              []GetAttributesResponseAttributeValue `json:"values"`
	IsMultipleSelection bool                                  `json:"is_multiple_selection"`
	IsCustomizable      bool                                  `json:"is_customizable"`
}

type GetAttributesResponseAttributeValue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetRecommendCategoryRequest struct {
	ProductTitle string                             `json:"product_title"`
	Description  *string                            `json:"description"`
	Images       []GetRecommendCategoryRequestImage `json:"images"`
}

type GetRecommendCategoryRequestImage struct {
	Uri string `json:"uri"`
}

type GetRecommendCategoryResponse struct {
	LeafCategoryId string                                 `json:"leaf_category_id"`
	Categories     []GetRecommendCategoryResponseCategory `json:"categories"`
}

type GetRecommendCategoryResponseCategory struct {
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	Level              int      `json:"level"`
	IsLeaf             bool     `json:"is_leaf"`
	PermissionStatuses []string `json:"permission_statuses"`
	/*
		VALUES:
		AVAILABLE
		INVITE_ONLY
	*/
}
