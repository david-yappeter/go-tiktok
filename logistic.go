package gotiktok

import (
	"context"
	"fmt"
)

// GetWarehouseList get warehouse list.
func (c *Client) GetWarehouseList(ctx context.Context, p Param) (list WarehouseList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/logistics/%s/warehouses", c.version),
		param,
		&list,
	)
	return
}
