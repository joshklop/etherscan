package etherscan

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	key    string
	client *http.Client
	url    string
}

type Option func(*Client)

func WithURL(url string) Option {
	return func(c *Client) {
		c.url = url
	}
}

func New(key string, opts ...Option) *Client {
	c := &Client{
		key:    key,
		url:    "https://api.etherscan.io/api",
		client: &http.Client{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) buildQueryString(endpoint string, args map[string]string) string {
	base, err := url.Parse(c.url)
	if err != nil {
		panic("malformed URL")
	}

	params := url.Values{}
	for k, v := range args {
		params.Add(k, v)
	}
	base.RawQuery = params.Encode()

	return base.String()
}

func (c *Client) get(ctx context.Context, query map[string]string) (io.ReadCloser, error) {
	url := c.buildQueryString("", query)
	req, err := http.NewRequestWithContext(ctx, "", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, errors.New(resp.Status)
	}
	return resp.Body, nil
}

type abiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func (c *Client) ABI(ctx context.Context, contractAddress string) (json.RawMessage, error) {
	body, err := c.get(ctx, map[string]string{
		"module":  "contract",
		"action":  "getabi",
		"address": contractAddress,
		"apikey":  c.key,
	})
	if err != nil {
		return nil, err
	}
	defer body.Close()
	resultBytes, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var resultABI abiResponse
	if err := json.Unmarshal(resultBytes, &resultABI); err != nil {
		return nil, err
	}
	var abi json.RawMessage
	if err := json.Unmarshal([]byte(resultABI.Result), &abi); err != nil {
		return nil, err
	}
	return abi, nil
}
