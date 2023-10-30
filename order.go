package gotiktok

import (
	"context"
	"fmt"
	"strings"
)

// GetOrderDetail get order detail.
func (c *Client) GetOrderDetail(ctx context.Context, p Param, orderIds []string) (list OrderDetailList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}

	param.Set("ids", strings.Join(orderIds, ","))

	err = c.Get(
		ctx,
		fmt.Sprintf("/order/%s/orders", c.version),
		param,
		&list,
	)
	return
}
