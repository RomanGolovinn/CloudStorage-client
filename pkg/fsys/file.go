package fsys

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

type FileSystemObject interface {
	IsDirectory() bool
	GetSize() int64
	UpdateFile([]byte) error
	UpdateDir() error
	GetHash() string
}

type File struct {
	Name  string `json:"name"`
	Path  string `json:"path"` //путь хранится без имени
	Size  int64  `json:"size"`
	IsDir bool   `json:"is_dir"`
	Hash  string `json:"hash"`
}

func (f File) IsDirectory() bool {
	return f.IsDir
}

func (f File) GetSize() int64 {
	return f.Size
}

func (f File) GetHash() string {
	return f.Hash
}

func (f File) UpdateFile(content []byte) error {
	_, err := os.Stat(f.Path)
	if err == os.ErrNotExist {
		return f.create(content)
	}
	// Идея: сделать проверку на совпадение контента
	// и исправлять только несовпадающий контент
	err = os.Remove(f.Path)
	if err != nil {
		return fmt.Errorf("can not remove file")
	}
	return f.create(content)
}

func (f File) UpdateDir() error {
	return nil
}

func (f File) create(content []byte) error {
	_, err := os.Create(f.Path)
	if err != nil {
		return fmt.Errorf("can not create file %s", f.Path)
	}
	file, err := os.Open(f.Path)
	if err != nil {
		return fmt.Errorf("can not open file %s", f.Path)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		return fmt.Errorf("can not write file %s", f.Path)
	}
	return nil
}

func (f File) CalculateHash() (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(f.Path))
	if err != nil {
		return "", fmt.Errorf("can not hash file %s path", f.Path)
	}
	_, err = hasher.Write([]byte(f.Name))
	if err != nil {
		return "", fmt.Errorf("can not hash file %s name", f.Path)
	}
	_, err = hasher.Write([]byte(strconv.FormatInt(f.Size, 10)))
	if err != nil {
		return "", fmt.Errorf("can not hash file %s size", f.Path)
	}
	content, err := f.GetContent()
	if err != nil {
		return "", err
	}
	_, err = hasher.Write(content)
	if err != nil {
		return "", fmt.Errorf("can not hash file %s content", f.Path)
	}

	hash := hasher.Sum(nil)

	return hex.EncodeToString(hash), nil
}

func (f File) GetContent() ([]byte, error) {
	content, err := os.ReadFile(f.Path)
	if err != nil {
		return []byte{}, fmt.Errorf("can not read file %s", f.Path)
	}
	return content, nil
}
