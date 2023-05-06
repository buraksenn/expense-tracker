package spreadsheet

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

type Client struct {
	svc *sheets.Service
}

func NewClient(ctx context.Context) (*Client, error) {
	s, err := sheets.NewService(ctx)
	if err != nil {
		return nil, err
	}
	return &Client{svc: s}, nil

}

func (c *Client) WriteNum(spreadsheetID string, r Range, value int) error {
	valueRange := sheets.ValueRange{
		Values: [][]interface{}{{value}},
	}
	_, err := c.svc.Spreadsheets.Values.Update(spreadsheetID, r.String(), &valueRange).ValueInputOption("RAW").Do()
	return err
}
