package gotiktok

import (
	"context"
	"fmt"
	"net/url"
)

func (c *Client) GetAuthorizedShops(ctx context.Context, accessToken string, shopId *string) (list ShopList, err error) {

	param := url.Values{}
	param.Set("access_token", accessToken)
	if shopId != nil {
		param.Set("shop_id", *shopId)
	}
	
	err = c.Get(
		ctx,
		fmt.Sprintf("/authorization/%s/shops", c.version),
		param,
		&list,
	)
	return
}
