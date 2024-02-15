package gotiktok

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) UploadImage(ctx context.Context, p CommonParam, request UploadProductImageRequest) (response UploadProductImageResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	if request.UseCase != nil {
		param.Set("use_case", *request.UseCase)
	}

	path := fmt.Sprintf("/product/%s/images/upload", c.version)

	param = c.prepareParam(path, param)

	err = c.requestMultipart(
		ctx,
		http.MethodPost,
		c.opt.endpoint,
		path,
		param,
		map[string][]io.Reader{
			"data": {request.Data},
		},
		&response,
	)

	return
}
