package structs

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type File struct {
	FileInfo `json:"info"`
	Content  []byte `json:"content"`
}

type FileInfo struct {
	Name  string `json:"name"`
	Path  string `json:"path"` //путь хранится без имени
	Size  int64  `json:"size"`
	IsDir bool   `json:"is_dir"`
}

func (f File) IsDirectory() bool {
	return f.IsDir
}

func (f File) GetSize() int64 {
	return f.Size
}

func (f File) Update() error {
	_, err := os.Stat(f.Path)
	if err == os.ErrNotExist {
		return f.create()
	}
	// Идея: сделать проверку на совпадение контента
	// и исправлять только несовпадающий контент
	err = os.Remove(f.Path)
	if err != nil {
		return fmt.Errorf("can not remove file")
	}
	return f.create()
}

func (f File) create() error {
	_, err := os.Create(f.Path)
	if err != nil {
		return fmt.Errorf("can not create file %s", f.Path)
	}
	file, err := os.Open(f.Path)
	if err != nil {
		return fmt.Errorf("can not open file %s", f.Path)
	}
	defer file.Close()

	_, err = file.Write(f.Content)
	if err != nil {
		return fmt.Errorf("can not write file %s", f.Path)
	}
	return nil
}

func ParseFile(path string, wg *sync.WaitGroup) (File, error) {
	defer wg.Done()
	pathlist := strings.Split(path, "/")
	name := pathlist[len(pathlist)-1]
	info, err := os.Stat(path)
	if err != nil {
		return File{}, fmt.Errorf("file %s not found", path)
	}
	size := info.Size()
	isDir := info.IsDir()
	content, err := getContent(path)
	if err != nil {
		return File{}, err
	}

	return File{
		FileInfo: FileInfo{
			Name:  name,
			Path:  path,
			Size:  size,
			IsDir: isDir,
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
