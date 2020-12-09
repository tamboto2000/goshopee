package goshopee

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	// base url for Shopee API
	baseAPIURL = "https://shopee.co.id/api"
	// default User Agent. Change User Agent by calling Shopee.SetUserAgent
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0"
)

var defaultHeader = http.Header{
	"User-Agent":        {defaultUserAgent},
	"Accept":            {"application/json"},
	"Accept-Language":   {"en-US,en;q=0.5"},
	"X-Requested-With":  {"XMLHttpRequest"},
	"X-API-SOURCE":      {"pc"},
	"Connection":        {"keep-alive"},
	"X-Shopee-Language": {"id"},
}

// Shopee is a client to wrap Shopee web internal API
type Shopee struct {
	cookies   *sync.Map
	client    *http.Client
	userAgent string
	csrfToken string
}

// New initiate Shopee client
func New() *Shopee {
	return &Shopee{
		cookies:   new(sync.Map),
		client:    new(http.Client),
		userAgent: defaultUserAgent,
	}
}

// SetCookieStr set cookies to client
func (sh *Shopee) SetCookieStr(c string) {
	header := http.Header{}
	header.Add("Cookie", c)
	request := http.Request{Header: header}
	for _, c := range request.Cookies() {
		sh.cookies.Store(c.Name, c)
	}

	// get csrf token
	for _, c := range request.Cookies() {
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
	sh.cookies.Range(func(key, value interface{}) bool {
		req.AddCookie(value.(*http.Cookie))

		return true
	})

	resp, err := sh.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sh.mergeCookies(resp.Cookies())

	if resp.StatusCode > 200 {
		return nil, errors.New(string(raw))
	}

	return raw, nil
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
	sh.cookies.Range(func(key, value interface{}) bool {
		req.AddCookie(value.(*http.Cookie))

		return true
	})

	// add csrf token to header
	req.Header.Set("X-CSRFToken", sh.csrfToken)
	// add content type
	req.Header.Set("Content-Type", "application/json")
	// add referer
	req.Header.Set("Referer", referer)

	resp, err := sh.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sh.mergeCookies(resp.Cookies())

	if resp.StatusCode > 299 {
		if len(raw) == 0 {
			return nil, errors.New(resp.Status)
		}

		return nil, errors.New(string(raw))
	}

	return raw, nil
}

func (sh *Shopee) mergeCookies(newC []*http.Cookie) {
	// replace old cookies with new one
	sh.cookies.Range(func(key, value interface{}) bool {
		for _, c := range newC {
			if key.(string) == c.Name {
				sh.cookies.Delete(key)
				sh.cookies.Store(key, c)
			}
		}

		return true
	})

	// insert new cookies
	for _, c := range newC {
		isNew := true
		sh.cookies.Range(func(key, value interface{}) bool {
			if key.(string) == c.Name {
				isNew = false
				return false
			}

			return true
		})

		if isNew {
			sh.cookies.Store(c.Name, c)
		}
	}
}
