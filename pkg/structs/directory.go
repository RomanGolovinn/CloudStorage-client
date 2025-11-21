package structs

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"CloudStorage-client/pkg/interfaces"
)

type Directory struct {
	FileInfo
	Children []interfaces.FileSystemObject
}

type parseDirResult struct {
	Object interfaces.FileSystemObject
	Err    error
}

func (d Directory) IsDirectory() bool {
	return d.IsDir
}

func (d Directory) GetSize() int64 {
	return d.Size
}

func (d Directory) Update() error {
	_, err := os.Stat(d.Path)
	if err == os.ErrNotExist {
		return d.create()
	}

	for _, obj := range d.Children {
		_ = obj.Update()

	}

	return err
}

func (d Directory) create() error {
	err := os.MkdirAll(d.Path+d.Name, 0755)
	if err != nil {
		return fmt.Errorf("can not create dir %s", (d.Path + d.Name))
	}
	return nil
}

func ParseDir(path string, wg *sync.WaitGroup) (Directory, error) {
	defer wg.Done()
	entries, err := os.ReadDir(path)
	if err != nil {
		return Directory{}, fmt.Errorf("%s", err)
	}

	pathlist := strings.Split(path, "/")
	name := pathlist[len(pathlist)-1]
	var size int64

	if len(entries) == 0 {
		size = 0
		return Directory{
			FileInfo: FileInfo{
				Name:  name,
				Path:  path,
				Size:  size,
				IsDir: true,
			},
			Children: []interfaces.FileSystemObject{},
		}, nil
	}

	children := make([]interfaces.FileSystemObject, len(entries))

	var dwg sync.WaitGroup //группа горутин для рекурсивного вызова
	dwg.Add(len(entries))
	results := make(chan parseDirResult, len(entries))

	for _, entry := range entries {
		if entry.IsDir() {
			go func() {
				dir, err := ParseDir(path+"/"+entry.Name(), &dwg)
				results <- parseDirResult{
					Object: dir,
					Err:    err,
				}
			}()
		} else {
			go func() {
				file, err := ParseFile(path+"/"+entry.Name(), &dwg)
				results <- parseDirResult{
					Object: file,
					Err:    err,
				}
			}()
		}
	}

	dwg.Wait()
	close(results)

	index := 0
	for result := range results {
		size += result.Object.GetSize()
		if result.Err != nil {
			return Directory{}, result.Err
		}
		children[index] = result.Object
		index++
	}

	return Directory{
		FileInfo: FileInfo{
			Name:  name,
			Path:  path,
			Size:  size,
			IsDir: true,
		},
		Children: children,
	}, nil
}
