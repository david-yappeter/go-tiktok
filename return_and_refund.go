package gotiktok

import (
	"context"
	"fmt"
)

// GetShippingProviders get shipping providers.
func (c *Client) SearchReturns(ctx context.Context, p CommonParam, cp CursorPaginationParam, req SearchReturnRequest) (list ShippingProviderList, err error) {
	param, err := c.commonParam(p)
	if err != nil {
		return
	}

	if param, err = c.cursorPaginationParam(param, cp); err != nil {
		return
	}

	if err = c.validate.Struct(&req); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/return_refund/%s/returns/search", c.version),
		param,
		req,
		&list,
	)
	return
}
