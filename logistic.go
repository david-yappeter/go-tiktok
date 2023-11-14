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

// GetWarehouseDetail get warehouse detail.
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

// GetShippingProviders get shipping providers.
func (c *Client) GetShippingProviders(ctx context.Context, p Param, deliveryOptionId string) (list ShippingProviderList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/logistics/%s/delivery_options/%s/shipping_providers", c.version, deliveryOptionId),
		param,
		&list,
	)
	return
}
