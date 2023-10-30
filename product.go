package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetProductDetail(ctx context.Context, p Param, productID string) (product ProductData, err error) {

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
