package service

import (
	"net/http"

	"github.com/buraksenn/expense-tracker/client/spreadsheet"
	"github.com/buraksenn/expense-tracker/pkg/storage"
)

type Service struct {
	SpreadsheetService *spreadsheet.Client
	Storage            storage.Storage
	Router             http.Handler
}

func New() *Service {
	return &Service{}
}
