package drive

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"google.golang.org/api/drive/v3"
)

type Client interface {
	Upload(userID string, content []byte) error
	GetFileLinksByUserID(userID string, from time.Time) ([]string, error)
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
	fileID := fmt.Sprintf("%s_%s", userID, date.Format("2006_01_02T15_04_05Z07_00"))
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

func (c *DefaultClient) GetFileLinksByUserID(userID string, from time.Time) ([]string, error) {
	l, err := c.svc.Files.List().Q(fmt.Sprintf("name contains '%s' and createdTime > '%s'", userID, from.Format(time.RFC3339))).Do()

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
