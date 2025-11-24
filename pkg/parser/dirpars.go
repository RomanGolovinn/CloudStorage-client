package parser

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"CloudStorage-client/pkg/fsys"
)

type parseDirResult struct {
	Object fsys.FileSystemObject
	Err    error
}

func ParseDir(path string, wg *sync.WaitGroup) (fsys.Directory, error) {
	defer wg.Done()
	entries, err := os.ReadDir(path)
	if err != nil {
		return fsys.Directory{}, fmt.Errorf("%s", err)
	}

	pathlist := strings.Split(path, "/")
	name := pathlist[len(pathlist)-1]
	var size int64

	if len(entries) == 0 {
		size = 0
		return fsys.Directory{
			FileInfo: fsys.FileInfo{
				Name:  name,
				Path:  path,
				Size:  size,
				IsDir: true,
			},
			Children: []fsys.FileSystemObject{},
		}, nil
	}

	children := make([]fsys.FileSystemObject, len(entries))

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
			return fsys.Directory{}, result.Err
		}
		children[index] = result.Object
		index++
	}

	return fsys.Directory{
		FileInfo: fsys.FileInfo{
			Name:  name,
			Path:  path,
			Size:  size,
			IsDir: true,
		},
		Children: children,
	}, nil
}
