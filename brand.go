package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetBrands(ctx context.Context, p CommonParam, cp CursorPaginationParam, request GetBrandsRequest) (response GetBrandsResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	if param, err = c.cursorPaginationParam(param, cp); err != nil {
		return
	}

	if err = c.validate.Struct(&request); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/brands", c.version),
		param,
		request,
		&response,
	)
	return
}
