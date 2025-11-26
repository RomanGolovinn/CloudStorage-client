package client

import (
	"CloudStorage-client/pkg/fsys"
	"encoding/json"
	"io"
)

type RespFile struct {
	File    fsys.File `json:"file"`
	Content []byte    `json:"content"`
}

func GetFile(ServPath string) (RespFile, error) {
	req := Request{
		Method: "Get",
		URL:    "http://localhost:8080/file", //localhost для тестирования
		Body:   "getfile{" + ServPath + "}",
	}
	resp, err := req.DoRequest()
	if err != nil {
		return RespFile{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespFile{}, err
	}

	file := RespFile{}
	err = json.Unmarshal(body, &file)
	if err != nil {
		return RespFile{}, err
	}

	// тут надо записать resp в File
	return file, nil
}
