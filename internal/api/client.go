package api

import (
	"CloudStorage-client/pkg/common"
	"net/http"
)

type Client struct {
	Token      string
	httpClient http.Client
	baseURL    string
}

func (c *Client) doRequest() {
	// doRequest выполняет HTTP-запрос,
	// добавляет заголовок авторизации и декодирует ответ.
}

func (c *Client) RequestFileList() []common.FileInfo {
	// запрос на список всех файлов
}
