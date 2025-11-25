package fsys

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var UpdateQueue = []FileSystemObject{}

type Directory struct {
	Name     string `json:"name"`
	Path     string `json:"path"` //путь хранится без имени
	Size     int64  `json:"size"`
	IsDir    bool   `json:"is_dir"`
	Hash     string `json:"hash"`
	Children []FileSystemObject
}

func (d Directory) IsDirectory() bool {
	return d.IsDir
}

func (d Directory) GetSize() int64 {
	return d.Size
}

func (d Directory) GetHash() string {
	return d.Hash
}

func (d Directory) UpdateDir() error {
	_, err := os.Stat(d.Path)
	if err == os.ErrNotExist {
		return d.create()
	}

	for _, obj := range d.Children {
		if obj.IsDirectory() {
			_ = obj.UpdateDir()
		} else {
			UpdateQueue = append(UpdateQueue, obj)
		}
	}

	return err
}

func (d Directory) UpdateFile(content []byte) error {
	return nil
}

func (d Directory) create() error {
	err := os.MkdirAll(d.Path+d.Name, 0755)
	if err != nil {
		return fmt.Errorf("can not create dir %s", (d.Path + d.Name))
	}
	return nil
}

func (d Directory) CalculateHash() (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(d.Path))
	if err != nil {
		return "", fmt.Errorf("can not hash dir %s path", d.Path)
	}
	_, err = hasher.Write([]byte(d.Name))
	if err != nil {
		return "", fmt.Errorf("can not hash dir %s name", d.Path)
	}
	_, err = hasher.Write([]byte(strconv.FormatInt(d.Size, 10)))
	if err != nil {
		return "", fmt.Errorf("can not hash dir %s size", d.Path)
	}
	_, err = hasher.Write([]byte(d.GetChildrenHash()))
	if err != nil {
		return "", fmt.Errorf("can not hash dir %s children", d.Path)
	}
	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash), nil
}

func (d Directory) GetChildrenHash() string {
	hashes := make([]string, len(d.Children))
	for i, child := range d.Children {
		h := child.GetHash()
		hashes[i] = h
	}
	return strings.Join(hashes, "")
}
