package client

import (
	"CloudStorage-client/pkg/structs"
)

type RespFile struct {
	File    structs.File
	Content []byte `json:"content"`
}

func GetFile(ServPath string) (structs.File, error) {
	req := Request{
		Method: "Get",
		URL:    "http://localhost:8080/file", //localhost для тестирования
		Body:   "getfile{" + ServPath + "}",
	}
	_, err := req.DoRequest()
	if err != nil {
		return structs.File{}, err
	}
	// тут надо записать resp в File
	return structs.File{}, nil
}
