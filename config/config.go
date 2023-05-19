package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type BaseConfig struct {
	Testing               string
	Debug                 string
	SecretKey             string
	Port                  string
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
	var (
		Testing                    = os.Getenv("Testing")
		Debug                      = os.Getenv("Debug")
		Port                       = os.Getenv("Port")
		SecretKey                  = os.Getenv("SecretKey")
		AWS_ACCCESS_KEY            = os.Getenv("AWS_ACCCESS_KEY")
		AWS_SECRET_KEY             = os.Getenv("AWS_DEFAULT_REGION")
		AWS_DEFAULT_REGION         = os.Getenv("AWS_DEFAULT_REGION")
		S3_Daily_Data              = os.Getenv("S3_Daily_Data")
		S3_Weekly_Data             = os.Getenv("S3_Weekly_Data")
		S3_Monthly_Data            = os.Getenv("S3_Monthly_Data")
		S3_Yearly_data             = os.Getenv("S3_Yearly_data")
		Test_S3_Daily_Data         = os.Getenv("Test_S3_Daily_Data")
		Test_S3_Weekly_Data        = os.Getenv("Test_S3_Weekly_Data")
		Test_S3_Monthly_Data       = os.Getenv("Test_S3_Monthly_Data")
		Test_S3_Yearly_data        = os.Getenv("Test_S3_Yearly_data")
		DynamoDB_UsersTable        = os.Getenv("DynamoDB_UsersTable")
		DynamoDB_ClientsTable      = os.Getenv("DynamoDB_ClientsTable")
		DynamoDB_DevicesTable      = os.Getenv("DynamoDB_DevicesTable")
		Test_DynamoDB_UsersTable   = os.Getenv("Test_DynamoDB_UsersTable")
		Test_DynamoDB_ClientsTable = os.Getenv("Test_DynamoDB_ClientsTable")
		Test_DynamoDB_DevicesTable = os.Getenv("Test_DynamoDB_DevicesTable")
		conf                       BaseConfig
	)

	switch env {
	case "dev":
		conf.Testing = Testing
		conf.S3_Daily_Data = Test_S3_Daily_Data
		conf.S3_Weekly_Data = Test_S3_Weekly_Data
		conf.S3_Monthly_Data = Test_S3_Monthly_Data
		conf.S3_Yearly_data = Test_S3_Yearly_data
		conf.DynamoDB_UsersTable = Test_DynamoDB_UsersTable
		conf.DynamoDB_ClientsTable = Test_DynamoDB_ClientsTable
		conf.DynamoDB_DevicesTable = Test_DynamoDB_DevicesTable

	case "prod":
		conf.S3_Daily_Data = S3_Daily_Data
		conf.S3_Weekly_Data = S3_Weekly_Data
		conf.S3_Monthly_Data = S3_Monthly_Data
		conf.S3_Yearly_data = S3_Yearly_data
		conf.DynamoDB_UsersTable = DynamoDB_UsersTable
		conf.DynamoDB_ClientsTable = DynamoDB_ClientsTable
		conf.DynamoDB_DevicesTable = DynamoDB_DevicesTable

	default:
		conf.Testing = Testing
		conf.S3_Daily_Data = Test_S3_Daily_Data
		conf.S3_Weekly_Data = Test_S3_Weekly_Data
		conf.S3_Monthly_Data = Test_S3_Monthly_Data
		conf.S3_Yearly_data = Test_S3_Yearly_data
		conf.DynamoDB_UsersTable = Test_DynamoDB_UsersTable
		conf.DynamoDB_ClientsTable = Test_DynamoDB_ClientsTable
		conf.DynamoDB_DevicesTable = Test_DynamoDB_DevicesTable
	}

	conf.Debug = Debug
	conf.SecretKey = SecretKey
	conf.Port = Port
	conf.AWS_ACCCESS_KEY = AWS_ACCCESS_KEY
	conf.AWS_SECRET_KEY = AWS_SECRET_KEY
	conf.AWS_DEFAULT_REGION = AWS_DEFAULT_REGION

	return &conf
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
