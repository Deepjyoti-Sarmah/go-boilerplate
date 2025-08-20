// Package handler
package handler

import (
	"github.com/deepjyoti-sarmah/go-boilerplate/internal/server"
	"github.com/deepjyoti-sarmah/go-boilerplate/internal/service"
)

type Handlers struct {
	Heath   *HeahtHander
	OpenAPI *OpenAPIHandler
}

func NewHandlers(s *server.Server, service *service.Service) *Handlers {
	return &Handlers{
		Heath:   NewHealthHandler(s),
		OpenAPI: NewOpenAPIHandler(s),
	}
}
