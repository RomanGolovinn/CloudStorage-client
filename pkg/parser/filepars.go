package parser

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"CloudStorage-client/pkg/fsys"
)

func ParseFile(path string, wg *sync.WaitGroup) (fsys.File, error) {
	defer wg.Done()
	pathlist := strings.Split(path, "/")
	name := pathlist[len(pathlist)-1]
	info, err := os.Stat(path)
	if err != nil {
		return fsys.File{}, fmt.Errorf("file %s not found", path)
	}
	size := info.Size()
	isDir := info.IsDir()
	content, err := getFileContent(path)
	if err != nil {
		return fsys.File{}, err
	}

	return fsys.File{
		FileInfo: fsys.FileInfo{
			Name:  name,
			Path:  path,
			Size:  size,
			IsDir: isDir,
		},
		Content: content,
	}, nil
}

func getFileContent(path string) ([]byte, error) {
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
