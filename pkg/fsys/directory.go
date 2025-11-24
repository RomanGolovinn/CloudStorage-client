package fsys

import (
	"fmt"
	"os"
)

type Directory struct {
	FileInfo
	Children []FileSystemObject
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
