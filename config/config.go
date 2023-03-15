package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Testing               = os.Getenv("Testing")
	Debug                 = os.Getenv("Debug")
	ENV                   = os.Getenv("ENV")
	Port                  = os.Getenv("Port")
	SecretKey             = os.Getenv("SecretKey")
	AWS_ACCCESS_KEY       = os.Getenv("AWS_ACCCESS_KEY")
	AWS_SECRET_KEY        = os.Getenv("AWS_DEFAULT_REGION")
	AWS_DEFAULT_REGION    = os.Getenv("AWS_DEFAULT_REGION")
	S3_Daily_Data         = os.Getenv("S3_Daily_Data")
	S3_Weekly_Data        = os.Getenv("S3_Weekly_Data")
	S3_Monthly_Data       = os.Getenv("S3_Monthly_Data")
	S3_Yearly_data        = os.Getenv("S3_Yearly_data")
	DynamoDB_UsersTable   = os.Getenv("DynamoDB_UsersTable")
	DynamoDB_ClientsTable = os.Getenv("DynamoDB_ClientsTable")
	DynamoDB_DevicesTable = os.Getenv("DynamoDB_DevicesTable")
)

type BaseConfig struct {
	Testing               string
	Debug                 string
	ENV                   string
	Port                  string
	SecretKey             string
	AWS_ACCCESS_KEY       string
	AWS_SECRET_KEY        string
	AWS_DEFAULT_REGION    string
	S3_Daily_Data         string
	S3_Weekly_Data        string
	S3_Monthly_Data       string
	S3_Yearly_data        string
	DynamoDB_UsersTable   string
	DynamoDB_ClientsTable string
	DynamoDB_DevicesTable string
}

func AppConfig(env string) *BaseConfig {
	if env == "Dev" {
		return &BaseConfig{
			Testing:               Testing,
			Debug:                 Debug,
			ENV:                   ENV,
			Port:                  Port,
			SecretKey:             SecretKey,
			AWS_ACCCESS_KEY:       AWS_ACCCESS_KEY,
			AWS_SECRET_KEY:        AWS_SECRET_KEY,
			AWS_DEFAULT_REGION:    AWS_DEFAULT_REGION,
			S3_Daily_Data:         S3_Daily_Data,
			S3_Weekly_Data:        S3_Weekly_Data,
			S3_Monthly_Data:       S3_Monthly_Data,
			S3_Yearly_data:        S3_Yearly_data,
			DynamoDB_UsersTable:   DynamoDB_UsersTable,
			DynamoDB_ClientsTable: DynamoDB_ClientsTable,
			DynamoDB_DevicesTable: DynamoDB_DevicesTable,
		}
	}
	return nil
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
