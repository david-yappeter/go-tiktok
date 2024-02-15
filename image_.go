package gotiktok

import "mime/multipart"

type UploadProductImageRequest struct {
	Data    multipart.File `json:"file"`
	UseCase *string        `json:"use_case"`
	/*
		Valid Values:
		MAIN_IMAGE
		ATTRIBUTE_IMAGE
		DESCRIPTION_IMAGE
		CERTIFICATION_IMAGE
		SIZE_CHART_IMAGE
	*/
}

type UploadProductImageResponse struct {
	Uri     string `json:"uri"`
	Url     string `json:"url"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	UseCase string `json:"use_case"`
}
