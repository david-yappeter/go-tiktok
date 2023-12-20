package gotiktok

import (
	"context"
	"fmt"
)

// GetTracking get tracking.
func (c *Client) GetTracking(ctx context.Context, p CommonParam, orderId string) (list TrackingList, err error) {
	param, err := c.commonParam(p)
	if err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/fulfillment/%s/orders/%s/tracking", c.version, orderId),
		param,
		&list,
	)
	return
}

// GetTracking get tracking.
func (c *Client) GetPackageDetail(ctx context.Context, p CommonParam, packageId string) (pack Package, err error) {
	param, err := c.commonParam(p)
	if err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/fulfillment/%s/packages/%s", c.version, packageId),
		param,
		&pack,
	)
	return
}
