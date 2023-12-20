package gotiktok

type Tracking struct {
	Description      string `json:"description"`
	UpdateTimeMillis int64  `json:"update_time_millis"`
}

type TrackingList struct {
	Trackings []Tracking `json:"trackings"`
}

type PackageOrderSku struct {
	Id       string `json:"id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type PackageOrder struct {
	Id   string            `json:"id"`
	Skus []PackageOrderSku `json:"skus"`
}

type Package struct {
	Id     string            `json:"package_id"`
	Status string            `json:"package_status"`
	Orders []PackageOrderSku `json:"orders"`
}
