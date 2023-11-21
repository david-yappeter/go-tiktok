package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) SearchProducts(ctx context.Context, p CommonParam, cp CursorPaginationParam, request SearchProductRequest) (productList ProductList, err error) {
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
		fmt.Sprintf("/product/%s/products/search", c.version),
		param,
		request,
		&productList,
	)
	return
}

func (c *Client) GetProductDetail(ctx context.Context, p CommonParam, productID string) (product ProductDetailData, err error) {

	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/products/%s", c.version, productID), param,
		&product,
	)
	return
}
