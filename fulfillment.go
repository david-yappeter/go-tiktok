package gotiktok

import (
	"context"
	"fmt"
)

// GetTracking get tracking.
func (c *Client) GetTracking(ctx context.Context, p Param, orderId string) (list TrackingList, err error) {
	param, err := c.params(p)
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
