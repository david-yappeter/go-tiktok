package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) SearchProducts(ctx context.Context,  p Param, request SearchProductRequest) (productList ProductList, err error) {
	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}

	if err = c.validate.Struct(&request); err != nil {
		return
	}

	param.Set("page_size", "10")

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products/search", c.version), 
		param,
		request,
		&productList,
	)
	return
}

func (c *Client) GetProductDetail(ctx context.Context, p Param, productID string) (product ProductDetailData, err error) {

	var param url.Values
	if param, err = c.params(p); err != nil {
		return
	}
	
	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/products/%s", c.version, productID), param,
		&product,
	)
	return
}
