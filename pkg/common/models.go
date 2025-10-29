package common

import (
	"time"
)

type FileInfo struct {
	Size           int64     `json:"size"`
	Name           string    `json:"name"`
	Path           string    `json:"path"`
	Hash           string    `json:"hash"`
	LastModified   time.Time `json:"last_modified"`
	ContentType    string    `json:"content_type"`
	CompressedType int       `json:"compressed"`
	EncryptedType  int       `json:"encrypted"`
}

func (f FileInfo) IsNewerThan(other FileInfo) bool {
	return f.LastModified.After(other.LastModified)
}

func (f FileInfo) HasChanged(other FileInfo) bool {
	return f.Size != other.Size || f.Hash != other.Hash || f.LastModified != other.LastModified
}

type DirectoryInfo struct {
	Files       []FileInfo      `json:"files"`
	Directories []DirectoryInfo `json:"directories"`
	Path        string          `json:"path"`
}

func (d DirectoryInfo) IsEmpty() bool {
	return (len(d.Files) + len(d.Directories)) == 0
}

type ServerInfo struct {
	Hostname    string `json:"hostname"`
	IP          string `json:"ip"`
	Environment string `json:"environment"`
}

type AuthRquest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      User      `json:"user"`
}

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	StorageQuota int64  `json:"storage_quota"`
	StorageUsed  int64  `json:"storage_used"`
}
