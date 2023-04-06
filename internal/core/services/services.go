/*
Package name : services
Filename : services.go
Description : Host code for business logic implementation
*/
package services

import (
	"errors"

	"github.com/AntonyIS/vision/internal/core/domain"
	"github.com/AntonyIS/vision/internal/core/ports"
)

var (
	ErrVisionNotFound      = errors.New("entity not found")
	ErrVisionInvalidEntity = errors.New("invalid entity")
	ErrInternalServerError = errors.New("internal server error")
)

type appService struct {
	repo ports.VisionRepository
}

func NewService(repo ports.VisionRepository) appService {
	return appService{
		repo: repo,
	}
}

func (a appService) CreateUser(user *domain.User) (*domain.User, error) {
	user.HashPassword()
	return a.repo.CreateUser(user)
}

func (a appService) ReadUser(user_id string) (*domain.User, error) {
	return a.repo.ReadUser(user_id)
}

func (a appService) ReadUsers() ([]*domain.User, error) {
	return a.repo.ReadUsers()
}

func (a appService) UpdateUser(user *domain.User) (*domain.User, error) {
	return a.repo.UpdateUser(user)
}

func (a appService) DeleteUser(user_id string) error {
	return a.repo.DeleteUser(user_id)
}

func (a appService) CreateDevice(device *domain.Device) (*domain.Device, error) {
	return a.repo.CreateDevice(device)
}

func (a appService) ReadDevice(device_id string) (*domain.Device, error) {
	return a.repo.ReadDevice(device_id)
}

func (a appService) ReadDevices() ([]*domain.Device, error) {
	return a.repo.ReadDevices()
}

func (a appService) UpdateDevice(device *domain.Device) (*domain.Device, error) {
	return a.repo.UpdateDevice(device)
}

func (a appService) DeleteDevice(device_id string) error {
	return a.repo.DeleteDevice(device_id)
}

func (a appService) CreateClient(client *domain.Client) (*domain.Client, error) {
	return a.repo.CreateClient(client)
}

func (a appService) ReadClient(client_id string) (*domain.Client, error) {
	return a.repo.ReadClient(client_id)
}

func (a appService) ReadClients() ([]*domain.Client, error) {
	return a.repo.ReadClients()
}

func (a appService) UpdateClient(client *domain.Client) (*domain.Client, error) {
	return a.repo.UpdateClient(client)
}

func (a appService) DeleteClient(client_id string) error {
	return a.repo.DeleteClient(client_id)
}
