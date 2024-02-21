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

	if request.BrandName != nil {
		param.Set("brand_name", *request.BrandName)
	}
	if request.CategoryId != nil {
		param.Set("category_id", *request.CategoryId)
	}
	if request.IsAuthorized != nil {
		param.Set("is_authorized", fmt.Sprintf("%+v\n", *request.IsAuthorized))
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/brands", c.version),
		param,
		&response,
	)
	return
}
