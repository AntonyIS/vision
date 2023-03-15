/*
Package name : ports
Filename : ports.go
Description : Host code for application business logic
*/
package ports

import "github.com/AntonyIS/vision1.0/core/domain"

type UserService interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ReadUser(user_id string) (*domain.User, error)
	ReadUsers() ([]*domain.User, error)
	UpdateUser(user *domain.User) (*domain.User, error)
	DeleteUser(user_id string) error
}

type ClientService interface {
	CreateClient(client *domain.Client) (*domain.Client, error)
	ReadClient(client_id string) (*domain.Client, error)
	ReadClients() ([]*domain.Client, error)
	UpdateClient(client *domain.Client) (*domain.Client, error)
	DeleteClient(client_id string) error
}

type DeviceService interface {
	CreateDevice(device *domain.Device) (*domain.Device, error)
	ReadDevice(device_id string) (*domain.Device, error)
	ReadDevices() ([]*domain.Device, error)
	UpdateDevice(device *domain.Device) (*domain.Device, error)
	DeleteDevice(device_id string) error
}
