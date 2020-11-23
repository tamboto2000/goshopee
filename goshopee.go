package goshopee

import (
	"bytes"
	"encoding/json"
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
	"Accept":           {"application/json"},
	"X-Requested-With": {"XMLHttpRequest"},
	"X-API-SOURCE":     {"pc"},
	"Connection":       {"keep-alive"},
}

// Shopee is a client to wrap Shopee web internal API
type Shopee struct {
	cookies   []*http.Cookie
	client    *http.Client
	userAgent string
	csrfToken string
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

	// get csrf token
	for _, c := range sh.cookies {
		if c.Name == "csrftoken" {
			sh.csrfToken = c.Value
			break
		}
	}
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

	sh.cookies = mergeCookies(sh.cookies, resp.Cookies())

	return raw, nil
}

func mergeCookies(old, newC []*http.Cookie) []*http.Cookie {
	cookies := make([]*http.Cookie, 0)
	// replace old cookie with the new one
	for i, cOld := range old {
		for _, cNew := range newC {
			if cOld.Name == cNew.Name {
				old[i] = cNew

				break
			}
		}
	}

	// add new cookies
	for _, cNew := range newC {
		isNew := true
		for _, cOld := range old {
			if cOld.Name == cNew.Name {
				isNew = false
				break
			}
		}

		if isNew {
			cookies = append(cookies, cNew)
		}
	}

	cookies = append(cookies, old...)

	return cookies
}

// request Shopee API with method POST
func (sh *Shopee) post(path, referer string, body map[string]interface{}) ([]byte, error) {
	uri, err := url.Parse(baseAPIURL + path)
	if err != nil {
		return nil, err
	}

	bodyByts, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri.String(), bytes.NewBuffer(bodyByts))
	if err != nil {
		return nil, err
	}

	req.Header = defaultHeader
	for _, c := range sh.cookies {
		req.AddCookie(c)
	}

	// add csrf token to header
	req.Header.Add("X-CSRFToken", sh.csrfToken)
	// add content type
	req.Header.Add("Content-Type", "application/json")
	// add referer
	req.Header.Add("Referer", referer)

	resp, err := sh.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		if len(raw) == 0 {
			return nil, errors.New(resp.Status)
		}

		return nil, errors.New(string(raw))
	}

	sh.cookies = mergeCookies(sh.cookies, resp.Cookies())

	return raw, nil
}
