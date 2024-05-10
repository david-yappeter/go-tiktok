package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) SearchProducts(ctx context.Context, p CommonParam, cp CursorPaginationParam, request SearchProductRequest) (productList ProductList, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	if param, err = c.cursorPaginationParam(param, cp); err != nil {
		return
	}

	if err = c.validate.Struct(&request); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products/search", c.version),
		param,
		request,
		&productList,
	)
	return
}

func (c *Client) GetProductDetail(ctx context.Context, p CommonParam, productID string) (product ProductDetailData, err error) {

	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/products/%s", c.version, productID), param,
		&product,
	)
	return
}

func (c *Client) CreateProduct(ctx context.Context, p CommonParam, request CreateProductRequest) (response CreateProductResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products", c.version),
		param,
		request,
		&response,
	)
	return
}

func (c *Client) ListingCheckProduct(ctx context.Context, p CommonParam, request ListingCheckProductRequest) (response ListingCheckProductResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products/listing_check", c.version),
		param,
		request,
		&response,
	)
	return
}

func (c *Client) UpdateProductInventory(ctx context.Context, p CommonParam, productId string, request UpdateProductInventoryRequest) (response UpdateProductInventoryResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products/%s/inventory/update", c.version, productId),
		param,
		request,
		&response,
	)
	return
}

func (c *Client) UpdateProductPrice(ctx context.Context, p CommonParam, productId string, request UpdateProductPriceRequest) (response UpdateProductPriceResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products/%s/prices/update", c.version, productId),
		param,
		request,
		&response,
	)
	return
}

func (c *Client) ActivateProduct(ctx context.Context, p CommonParam, request ActivateProductRequest) (response UpdateProductPriceResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products/activate", c.version),
		param,
		request,
		&response,
	)
	return
}

func (c *Client) DeactivateProduct(ctx context.Context, p CommonParam, request DeactivateProductRequest) (response UpdateProductPriceResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Post(
		ctx,
		fmt.Sprintf("/product/%s/products/deactivate", c.version),
		param,
		request,
		&response,
	)
	return
}

func (c *Client) UpdateProduct(ctx context.Context, p CommonParam, productId string, request UpdateProductRequest) (response UpdateProductResponse, err error) {
	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Put(
		ctx,
		fmt.Sprintf("/product/%s/products/%s", c.version, productId),
		param,
		request,
		&response,
	)
	return
}
