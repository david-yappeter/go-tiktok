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

func (c *Client) GetWarehouseDeliveryOptions(ctx context.Context, p Param, warehouseId string) (list WarehouseDeliveryOptionList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/logistics/%s/warehouses/%s/delivery_options", c.version, warehouseId),
		param,
		&list,
	)
	return
}
