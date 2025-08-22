// Package service
package service

import (
	"github.com/deepjyoti-sarmah/go-boilerplate/internal/lib/job"
	"github.com/deepjyoti-sarmah/go-boilerplate/internal/repository"
	"github.com/deepjyoti-sarmah/go-boilerplate/internal/server"
)

type Service struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Service, error) {
	authService := NewAuthService(s)

	return &Service{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
