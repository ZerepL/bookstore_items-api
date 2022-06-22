package services

import (
	"github.com/ZerepL/bookstore_items-api/domain/items"
	"github.com/ZerepL/bookstore_utils/internal_errors"
	"net/http"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, internal_errors.RestErr)
	Get(string) (*items.Item, internal_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(items.Item) (*items.Item, internal_errors.RestErr) {
	return nil, internal_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}

func (s *itemsService) Get(string) (*items.Item, internal_errors.RestErr) {
	return nil, internal_errors.NewRestError("implement me", http.StatusNotImplemented, "not_implemented", nil)
}
