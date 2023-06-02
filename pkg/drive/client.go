package drive

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
)

type Client interface {
	Upload(userID string, content []byte) error
	GetFileLinksByUserID(userID string) ([]string, error)
}

type DefaultClient struct {
	svc *drive.Service
}

func NewClient(ctx context.Context) (*DefaultClient, error) {
	s, err := drive.NewService(ctx)
	if err != nil {
		return nil, err
	}
	return &DefaultClient{svc: s}, nil
}

func (c *DefaultClient) Upload(userID string, content []byte) error {
	date := time.Now()
	fileID := fmt.Sprintf("%d-%d-%s", date.Month(), date.Day(), uuid.NewString())
	_, err := c.svc.Files.Create(&drive.File{
		Name: fileID,
		Permissions: []*drive.Permission{
			{
				Type: "anyone",
			},
		},
	}).Media(bytes.NewReader(content)).Do()
	return err
}

func (c *DefaultClient) GetFileLinksByUserID(userID string) ([]string, error) {
	// get files by filtering with userID
	l, err := c.svc.Files.List().Q(fmt.Sprintf("name contains '%s'", userID)).Do()
	if err != nil {
		return nil, err
	}

	if l == nil || len(l.Files) == 0 {
		return nil, nil
	}

	var links []string
	for _, f := range l.Files {
		links = append(links, f.WebContentLink)
	}
	return links, nil

}
