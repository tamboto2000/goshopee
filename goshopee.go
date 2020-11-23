package goshopee

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// base url for Shopee API
	baseAPIURL = "https://shopee.co.id/api/v2"
	// default User Agent. Change User Agent by calling Shopee.SetUserAgent
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0"
)

var defaultHeader = http.Header{
	"User-Agent":       {defaultUserAgent},
	"Accept":           {"Accept"},
	"X-Requested-With": {"XMLHttpRequest"},
	"X-API-SOURCE":     {"pc"},
	"Connection":       {"keep-alive"},
}

// Shopee is a client to wrap Shopee web internal API
type Shopee struct {
	cookies   []*http.Cookie
	client    *http.Client
	userAgent string
}

// New initiate Shopee client
func New() *Shopee {
	return &Shopee{
		client:    new(http.Client),
		userAgent: defaultUserAgent,
	}
}

// SetCookieStr set cookies to client
func (sh *Shopee) SetCookieStr(c string) {
	header := http.Header{}
	header.Add("Cookie", c)
	request := http.Request{Header: header}
	sh.cookies = request.Cookies()
}

// SetProxy set proxy to client
func (sh *Shopee) SetProxy(prox string) error {
	uri, err := url.Parse(prox)
	if err != nil {
		return err
	}

	sh.client.Transport = &http.Transport{Proxy: http.ProxyURL(uri), TLSHandshakeTimeout: 20 * time.Second}

	return nil
}

// SetUserAgent set client User Agent
func (sh *Shopee) SetUserAgent(agent string) {
	sh.userAgent = agent
}

// request Shopee API with method GET
func (sh *Shopee) get(path string, param url.Values) ([]byte, error) {
	uri, err := url.Parse(baseAPIURL + path)
	if err != nil {
		return nil, err
	}

	if param != nil {
		uri.RawQuery = param.Encode()
	}

	req, err := http.NewRequest("GET", uri.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header = defaultHeader
	for _, c := range sh.cookies {
		req.AddCookie(c)
	}

	resp, err := sh.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 200 {
		return nil, errors.New(string(raw))
	}

	return raw, nil
}
