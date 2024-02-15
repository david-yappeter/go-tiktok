package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetCategories(ctx context.Context, p CommonParam, request GetCategoriesRequest) (response GetCategoryResponse, err error) {

	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	param.Set("locale", *request.Locale)

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/categories", c.version),
		param,
		&response,
	)
	return
}

func (c *Client) GetCategoryRules(ctx context.Context, p CommonParam, categoryId string) (response GetCategoryRulesResponse, err error) {

	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/categories/%s/rules", c.version, categoryId),
		param,
		&response,
	)
	return
}

func (c *Client) GetAttributes(ctx context.Context, p CommonParam, categoryId string, request GetAttributesRequest) (response GetAttributesResponse, err error) {

	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/categories/%s/attributes", c.version, categoryId),
		param,
		&response,
	)
	return
}

func (c *Client) GetRecommendCategory(ctx context.Context, p CommonParam, request GetRecommendCategoryRequest) (response GetRecommendCategoryResponse, err error) {

	var param url.Values
	if param, err = c.commonParam(p); err != nil {
		return
	}

	err = c.Get(
		ctx,
		fmt.Sprintf("/product/%s/categories/recommend", c.version),
		param,
		&response,
	)
	return
}
