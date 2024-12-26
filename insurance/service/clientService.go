package service

import (
	"github.com/jmechavez/restapi_go/insurance/domain"
	"github.com/jmechavez/restapi_go/insurance/errs"
)

type ClientService interface {
	GetAllClient(string) ([]domain.Client, *errs.AppError)
	FindName(string) (*domain.Client, *errs.AppError)
	JustFName(string) (*domain.Client, error)
}

type DefaultClientService struct {
	repo domain.ClientRepository
}

func (r DefaultClientService) GetAllClient(status string) ([]domain.Client, *errs.AppError) {
	if status == "active" {
		status = "true"
	} else if status == "inactive" {
		status = "false"
	} else {
		status = ""
	}
	return r.repo.FindAll(status)
}

func (r DefaultClientService) FindName(id string) (*domain.Client, *errs.AppError) {
	return r.repo.ByName(id)
}

func (r DefaultClientService) JustFName(id string) (*domain.Client, error) {
	return r.repo.JustName(id)
}

func NewClientService(repository domain.ClientRepository) DefaultClientService {
	return DefaultClientService{repository}
}
