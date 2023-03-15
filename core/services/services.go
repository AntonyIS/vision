/*
Package name : services
Filename : services.go
Description : Host code for business logic implementation
*/
package services

import (
	"errors"

	"github.com/AntonyIS/vision1.0/core/domain"
	"github.com/AntonyIS/vision1.0/core/ports"
)

var (
	ErrVisionNotFound      = errors.New("entity not found")
	ErrVisionInvalidEntity = errors.New("invalid entity")
	ErrInternalServerError = errors.New("internal server error")
)

type appService struct {
	repo *ports.VisionRepository
}

func NewService(repo *ports.VisionRepository) *appService {
	return &appService{
		repo: repo,
	}
}

func (a *appService) CreateUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (a *appService) ReadUser(user_id string) (*domain.User, error) {
	return nil, nil
}

func (a *appService) ReadUsers() ([]*domain.User, error) {
	return nil, nil
}

func (a *appService) UpdateUser(user *domain.User) (*domain.User, error) {
	return nil, nil
}

func (a *appService) DeleteUser(user_id string) error {
	return nil
}

func (a *appService) CreateDevice(device *domain.Device) (*domain.Device, error) {
	return nil, nil
}

func (a *appService) ReadDevice(device_id string) (*domain.Device, error) {
	return nil, nil
}

func (a *appService) ReadDevices() ([]*domain.Device, error) {
	return nil, nil
}

func (a *appService) UpdateDevice(device *domain.Device) (*domain.Device, error) {
	return nil, nil
}

func (a *appService) DeleteDevice(device_id string) error {
	return nil
}

func (a *appService) CreateClient(client *domain.Client) (*domain.Client, error) {
	return nil, nil
}

func (a *appService) ReadClient(client_id string) (*domain.Client, error) {
	return nil, nil
}

func (a *appService) ReadClients() ([]*domain.Client, error) {
	return nil, nil
}

func (a *appService) UpdateClient(client *domain.Client) (*domain.Client, error) {
	return nil, nil
}

func (a *appService) DeleteClient(client_id string) error {
	return nil
}
