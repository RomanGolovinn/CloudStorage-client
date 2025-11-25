package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Method string
	URL    string
	Body   interface{}
}

func (req *Request) DoRequest() (*http.Response, error) {
	jsonData, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(jsonData)

	client := &http.Client{}

	request, err := http.NewRequest(req.Method, req.URL, reqBody)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
