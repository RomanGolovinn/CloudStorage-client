package structs

import (
	"fmt"
	"os"
	"strings"
)

type File struct {
	FileInfo FileInfo `json:"info"`
	Content  []byte   `json:"content"`
}

type FileInfo struct {
	Name string `json:"name"`
	Path string `json:"path"` //путь хранится без имени
	Size int64  `json:"size"`
}

func ParseFile(path string) (File, error) {
	pathlist := strings.Split(path, "/")
	name := pathlist[len(pathlist)-1]
	info, err := os.Stat(path)
	if err != nil {
		return File{}, fmt.Errorf("file %s not found", path)
	}
	size := info.Size()
	content, err := getContent(path)
	if err != nil {
		return File{}, err
	}

	return File{
		FileInfo: FileInfo{
			Name: name,
			Path: path,
			Size: size,
		},
		Content: content,
	}, nil

}

func getContent(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte{}, fmt.Errorf("can not open filr %s", path)
	}
	defer file.Close()

	content, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("can not read file %s", path)
	}
	return content, nil
}
