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

type AppService struct {
	repo ports.VisionRepository
}

func NewService(repo ports.VisionRepository) AppService {
	return AppService{
		repo: repo,
	}
}

func (a AppService) CreateUser(user *domain.User) (*domain.User, error) {
	user.HashPassword()
	return a.repo.CreateUser(user)
}

func (a AppService) ReadUser(user_id string) (*domain.User, error) {
	return a.repo.ReadUser(user_id)
}

func (a AppService) ReadUsers() ([]*domain.User, error) {
	return a.repo.ReadUsers()
}

func (a AppService) UpdateUser(user *domain.User) (*domain.User, error) {
	return a.repo.UpdateUser(user)
}

func (a AppService) DeleteUser(user_id string) error {
	return a.repo.DeleteUser(user_id)
}

func (a AppService) CreateDevice(device *domain.Device) (*domain.Device, error) {
	return a.repo.CreateDevice(device)
}

func (a AppService) ReadDevice(device_id string) (*domain.Device, error) {
	return a.repo.ReadDevice(device_id)
}

func (a AppService) ReadDevices() ([]*domain.Device, error) {
	return a.repo.ReadDevices()
}

func (a AppService) UpdateDevice(device *domain.Device) (*domain.Device, error) {
	return a.repo.UpdateDevice(device)
}

func (a AppService) DeleteDevice(device_id string) error {
	return a.repo.DeleteDevice(device_id)
}

func (a AppService) CreateClient(client *domain.Client) (*domain.Client, error) {
	return a.repo.CreateClient(client)
}

func (a AppService) ReadClient(client_id string) (*domain.Client, error) {
	return a.repo.ReadClient(client_id)
}

func (a AppService) ReadClients() ([]*domain.Client, error) {
	return a.repo.ReadClients()
}

func (a AppService) UpdateClient(client *domain.Client) (*domain.Client, error) {
	return a.repo.UpdateClient(client)
}

func (a AppService) DeleteClient(client_id string) error {
	return a.repo.DeleteClient(client_id)
}
