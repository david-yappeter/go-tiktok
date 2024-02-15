package gotiktok

type GetBrandsRequest struct {
	CategoryId   *string `json:"category_id"`
	IsAuthorized *bool   `json:"is_authorized"`
	BrandName    *string `json:"brand_name"`
}

type GetBrandsResponse struct {
	Brands        []GetBrandsResponseBrand `json:"brands"`
	TotalCount    int                      `json:"total_count"`
	NextPageToken string                   `json:"next_page_token"`
}

type GetBrandsResponseBrand struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	AuthorizedStatus string `json:"authorized_status"`
	IsT1Brand        bool   `json:"is_t1_brand"`
	BrandStatus      string `json:"brand_status"`
	/*
		Values:
		AVAILABLE
		UNAVAILABLE
	*/
}
