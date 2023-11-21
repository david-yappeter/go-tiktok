package gotiktok

import (
	"context"
	"fmt"
)

// GetShippingProviders get shipping providers.
func (c *Client) SearchReturns(ctx context.Context, p Param, req SearchReturnRequest, nextPageToken string, pageSize string) (list ShippingProviderList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}

	if err = c.validate.Struct(&req); err != nil {
		return
	}

	param.Set("next_page_token", nextPageToken)
	param.Set("page_size", pageSize)

	err = c.Post(
		ctx,
		fmt.Sprintf("/return_refund/%s/returns/search", c.version),
		param,
		req,
		&list,
	)
	return
}

