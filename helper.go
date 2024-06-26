package gotiktok

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"mime"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

func sign(req *http.Request, secret string) string {
	queries := req.URL.Query()

	// extract all query parameters excluding sign and access_token
	keys := make([]string, len(queries))
	idx := 0
	for k := range queries {
		// params except 'sign' & 'access_token'
		if k != "sign" && k != "access_token" {
			keys[idx] = k
			idx++
		}
	}

	// reorder the parameters' key in alphabetical order
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Concatenate all the parameters in the format of {key}{value}
	input := ""
	for _, key := range keys {
		input = input + key + queries.Get(key)
	}

	// append the request path
	input = req.URL.Path + input

	// if the request header Content-type is not multipart/form-data, append body to the end
	mediaType, _, _ := mime.ParseMediaType(req.Header.Get("Content-type"))
	if mediaType != "multipart/form-data" {
		if req.Body == nil {
			req.Body = http.NoBody
		} else {
			body, _ := io.ReadAll(req.Body)
			input = input + string(body)

			req.Body.Close()
			// reset body after reading from the original
			req.Body = io.NopCloser(bytes.NewReader(body))
		}
	}

	// wrap the string generated in step 5 with the App secret
	input = secret + input + secret

	return generateSHA256(input, secret)
}

func generateSHA256(input, secret string) string {
	// encode the digest byte stream in hexadecimal and use sha256 to generate sign with salt(secret)
	h := hmac.New(sha256.New, []byte(secret))

	if _, err := h.Write([]byte(input)); err != nil {
		// TODO error log
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}

func safeGet(param url.Values, key string) string {
	if param == nil {
		return ""
	}
	vs, ok := param[key]
	if !ok || len(vs) == 0 {
		return ""
	}
	return vs[0]
}

func CheckEmpty(vals ...string) bool {
	for _, v := range vals {
		if v == "" {
			return true
		}
	}
	return false
}

type CommonParam struct {
	AccessToken string `validate:"required"`
	ShopCipher  string `validate:"required"`
	ShopId      string `validate:"required"`
}

func (c *Client) commonParam(p CommonParam) (param url.Values, err error) {
	err = c.validate.Struct(p)
	if err != nil {
		return
	}
	param = url.Values{}
	param.Set("access_token", p.AccessToken)
	param.Set("shop_cipher", p.ShopCipher)
	param.Set("shop_id", p.ShopId)
	return
}

type CursorPaginationParam struct {
	NextPageToken *string
	PageSize      int     `validate:"min=10,max=100"`
	SortField     *string `validate:"omitempty,oneof=create_time update_time"`
	SortOrder     *string `validate:"omitempty,oneof=ASC DESC"`
}

func (c *Client) cursorPaginationParam(parent url.Values, cp CursorPaginationParam) (param url.Values, err error) {
	err = c.validate.Struct(cp)
	if err != nil {
		return
	}

	param = url.Values{}
	if parent != nil {
		param = parent
	}

	param.Set("page_size", strconv.Itoa(cp.PageSize))
	if cp.NextPageToken != nil {
		param.Set("page_token", *cp.NextPageToken)
	}
	if cp.SortField != nil {
		param.Set("sort_field", *cp.SortField)
	}
	if cp.SortOrder != nil {
		param.Set("sort_order", *cp.SortOrder)
	}
	return
}
