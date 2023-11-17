package gotiktok

type Tracking struct {
	Description      string `json:"description"`
	UpdateTimeMillis int64  `json:"update_time_millis"`
}

type TrackingList struct {
	Trackings []Tracking `json:"trackings"`
}
