/*
Package name : domain
Filename : domain.go
Description : Host code for application entities : Users, Devices, Company
*/
package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	Password  string `json:"password"`
	Image     string `json:"image"`
	Client    struct {
		CompanyID   string `json:"company_id"`
		CompanyName string `json:"company_name"`
	} `json:"client"`
	Devices []*struct {
		DeviceID   string `json:"device_id"`
		DeviceName string `json:"device_name"`
	} `json:"devices"`
	IsAdmin      bool `json:"is_admin"`
	IsSuperUser  bool `json:"is_super_user"`
	IsInstaller  bool `json:"is_installer"`
	IsClient     bool `json:"is_client"`
	IsManager    bool `json:"is_manager"`
	IsTechnician bool `json:"is_technician"`
}

type Device struct {
	DeviceID   string `json:"device_id"`
	DeviceName string `json:"device_name"`
	Image      string `json:"image"`
	Capacity   int64  `json:"capacity"`
	Client     struct {
		CompanyID   string `json:"company_id"`
		CompanyName string `json:"company_name"`
	} `json:"client"`
	Env       string `json:"env"`
	AppType   string `json:"app_type"`
	Latitude  int64  `json:"latitude"`
	Longitude int64  `json:"longitude"`
}

type Client struct {
	CompanyID   string             `json:"company_id"`
	CompanyName string             `json:"company_name"`
	Image       string             `json:"image"`
	Devices     map[string]*Device `json:"devices"`
	Users       map[string]*User   `json:"users"`
}

func (u User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
