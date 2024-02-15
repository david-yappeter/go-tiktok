package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetCategories(ctx context.Context, p CommonParam, request GetCategoriesRequest) (response GetCategoryResponse, err error) {

	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	param.Set("locale", *request.Locale)

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/categories", c.version),
		param,
		&response,
	)
	return
}
