package gotiktok

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	// ServiceBaseURL for service related api call.
	ServiceBaseURL = "https://services.tiktokshop.com"
	// AuthBaseURL for auth related api call.
	AuthBaseURL = "https://auth.tiktok-shops.com"
	// APIBaseURL other apis except auth.
	APIBaseURL = "https://open-api.tiktokglobalshop.com"
)

// Timestamp mock this if you want to repeatable request.
var Timestamp func() string = func() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

// Client for tiktok
type Client struct {
	appKey    string
	appSecret string
	opt       *option
	validate  *validator.Validate
	endpoint  string
	version   string
}

// New for create a client.
func New(appKey, appSecret, version string, opts ...Option) (c *Client, err error) {
	if CheckEmpty(appKey, appSecret) {
		err = ErrAppInfoEmpty
		return
	}
	opt := defaultOpt()
	for _, fn := range opts {
		fn(opt)
	}
	c = &Client{
		appKey:    appKey,
		appSecret: appSecret,
		opt:       opt,
		validate:  validator.New(),
		endpoint:  opt.endpoint,
		version:   version,
	}
	return
}

// Get request for TikTok requests.
// Note: Timestamp, appkey and signature will auto-management by action.
func (c *Client) Get(ctx context.Context, path string, param url.Values, resp interface{}) (err error) {
	param = c.prepareParam(path, param)
	err = c.requestJson(ctx, http.MethodGet, c.opt.endpoint, path, param, nil, resp)
	return
}

// Post request for TikTok requests.
// Note: Timestamp, appkey and signature will auto-management by action.
func (c *Client) Post(ctx context.Context, path string, param url.Values, body interface{}, resp interface{}) (err error) {
	param = c.prepareParam(path, param)
	r := c.prepareBody(body)
	err = c.requestJson(ctx, http.MethodPost, c.opt.endpoint, path, param, r, resp)
	return
}

// Put request for TikTok requests.
// Note: Timestamp, appkey and signature will auto-management by action.
func (c *Client) Put(ctx context.Context, path string, param url.Values, body interface{}, resp interface{}) (err error) {
	param = c.prepareParam(path, param)
	r := c.prepareBody(body)
	err = c.requestJson(ctx, http.MethodPut, c.opt.endpoint, path, param, r, resp)
	return
}

// Delete request for TikTok requests. I don't known why there is body in delete request.
// Note: Timestamp, appkey and signature will auto-management by action.
func (c *Client) Delete(ctx context.Context, path string, param url.Values, body interface{}, resp interface{}) (err error) {
	param = c.prepareParam(path, param)
	var r io.Reader
	if body != nil {
		r = c.prepareBody(body)
	}
	err = c.requestJson(ctx, http.MethodDelete, c.opt.endpoint, path, param, r, resp)
	return
}

func (c *Client) prepareParam(path string, param url.Values) url.Values {
	ak := safeGet(param, "access_token")
	if ak != "" {
		param.Del("access_token")
	}
	param.Set("app_key", c.appKey)
	timestamp := Timestamp()
	param.Set("timestamp", timestamp)
	param.Set("access_token", ak)
	return param
}

func (c *Client) prepareBody(body interface{}) (buf io.Reader) {
	b, _ := json.Marshal(body)
	c.opt.logger.Printf("body: %s", string(b))
	buf = bytes.NewReader(b)
	return
}

func (c *Client) requestJson(ctx context.Context, method, base, path string, param url.Values, r io.Reader, body interface{}) (err error) {
	var req *http.Request
	target := base + path + "?" + param.Encode()
	c.opt.logger.Printf("%s %s", method, target)
	req, err = http.NewRequest(method, target, r)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Go-tiktok/v0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-tts-access-token", safeGet(param, "access_token"))
	req = req.WithContext(ctx)

	// append sign to the query param
	param.Set("sign", sign(req, c.appSecret))
	req.URL.RawQuery = param.Encode()

	resp, err := c.opt.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var res Response
	defer func() {
		c.opt.logger.Printf("request_id=%s resp=%s", res.RequestID, string(b))
	}()
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}
	if res.Code != 0 {
		err = &APIError{
			Code:      res.Code,
			Message:   res.Message,
			RequestID: res.RequestID,
		}
		return
	}

	if body != nil {
		err = json.Unmarshal(res.Data, body)
	}
	return
}

func (c *Client) requestMultipart(ctx context.Context, method, base, path string, param url.Values, fileParams map[string][]io.Reader, body interface{}) (err error) {
	var req *http.Request
	target := base + path + "?" + param.Encode()
	c.opt.logger.Printf("%s %s", method, target)

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	for paramName, files := range fileParams {
		for i, file := range files {
			fileWriter, err := bodyWriter.CreateFormFile(paramName+strconv.Itoa(i), filepath.Base(paramName+strconv.Itoa(i)))
			if err != nil {
				return err
			}
			_, err = io.Copy(fileWriter, file)
			if err != nil {
				return err
			}
		}
	}

	for key, values := range param {
		for _, value := range values {
			_ = bodyWriter.WriteField(key, value)
		}
	}

	bodyWriter.Close()

	req, err = http.NewRequest(method, target, bodyBuf)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "Go-tiktok/v0")
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.Header.Set("x-tts-access-token", safeGet(param, "access_token"))
	req = req.WithContext(ctx)

	// append sign to the query param
	param.Set("sign", sign(req, c.appSecret))
	req.URL.RawQuery = param.Encode()

	resp, err := c.opt.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var res Response
	defer func() {
		c.opt.logger.Printf("request_id=%s resp=%s", res.RequestID, string(b))
	}()
	err = json.Unmarshal(b, &res)
	if err != nil {
		return
	}
	if res.Code != 0 {
		err = &APIError{
			Code:      res.Code,
			Message:   res.Message,
			RequestID: res.RequestID,
		}
		return
	}

	if body != nil {
		err = json.Unmarshal(res.Data, body)
	}
	return
}
