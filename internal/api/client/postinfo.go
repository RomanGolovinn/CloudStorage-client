package client

import (
	"encoding/json"
)

func PostFile(file RespFile) error {
	body, err := json.Marshal(file)
	if err != nil {
		return err
	}
	req := Request{
		Method: "Post",
		URL:    "http://localhost:8080/file", //localhost для тестирования
		Body:   body,
	}
	resp, err := req.DoRequest()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
