package gotiktok

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GenerateAuthURL(
	serviceId string,
	state string,
) string {
	return fmt.Sprintf(
		ServiceBaseURL+"/open/authorize?service_id=%s&state=%s",
		serviceId, state,
	)
}

func (c *Client) GetAccessToken(ctx context.Context, code string) (resp AccessTokenResponse, err error) {
	param := url.Values{}
	param.Set("app_key", c.appKey)
	param.Set("app_secret", c.appSecret)
	param.Set("auth_code", code)
	param.Set("grant_type", "authorized_code")
	err = c.requestJson(
		ctx,
		http.MethodGet,
		AuthBaseURL,
		"/api/v2/token/get",
		param,
		nil,
		&resp,
	)
	return
}

func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (resp AccessTokenResponse, err error) {
	param := url.Values{}
	param.Set("app_key", c.appKey)
	param.Set("app_secret", c.appSecret)
	param.Set("refresh_token", refreshToken)
	param.Set("grant_type", "refresh_token")
	err = c.requestJson(
		ctx,
		http.MethodGet,
		AuthBaseURL,
		"/api/v2/token/refresh",
		param,
		nil,
		&resp,
	)
	return
}
