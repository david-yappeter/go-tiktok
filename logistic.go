package gotiktok

import (
	"context"
	"fmt"
)

// GetWarehouseList get warehouse list.
func (c *Client) GetWarehouseList(ctx context.Context, p CommonParam) (list WarehouseList, err error) {
	param, err := c.commonParam(p)
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
func (c *Client) GetWarehouseDeliveryOptions(ctx context.Context, p CommonParam, warehouseId string) (list WarehouseDeliveryOptionList, err error) {
	param, err := c.commonParam(p)
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
func (c *Client) GetShippingProviders(ctx context.Context, p CommonParam, deliveryOptionId string) (list ShippingProviderList, err error) {
	param, err := c.commonParam(p)
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
