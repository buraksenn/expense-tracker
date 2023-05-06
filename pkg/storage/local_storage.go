package storage

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type LocalStorage struct {
	path string
}

func NewLocalStorage(p string) *LocalStorage {
	return &LocalStorage{
		path: p,
	}
}

func (l *LocalStorage) Upload(content []byte) error {
	date := time.Now()
	fileID := fmt.Sprintf("%d-%d-%s", date.Month(), date.Day(), uuid.NewString())
	return os.WriteFile(fmt.Sprintf("%s/%s", l.path, fileID), content, 0644)

}
