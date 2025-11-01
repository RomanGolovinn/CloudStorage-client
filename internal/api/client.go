package api

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	Token      string
	httpClient http.Client
	baseURL    string
}

func NewClient(baseURL, token string) *Client {
	return &Client{
		Token: token,
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		baseURL: baseURL,
	}
}

func (c *Client) doRequest(method string, url string, body interface{}) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Body = ioutil.NopCloser(bytes.NewBufferString(fmt.Sprintf("{%s: %s}", "body", body)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(req.Body.(*io.Reader).Size())))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
