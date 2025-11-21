package interfaces

type FileSystemObject interface {
	IsDirectory() bool
	GetSize() int64
	Update() error
}
