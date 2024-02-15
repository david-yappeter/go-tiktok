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
