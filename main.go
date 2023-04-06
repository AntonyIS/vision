package main

import (
	"flag"

	"github.com/AntonyIS/vision/config"
	"github.com/AntonyIS/vision/internal/adapters/database/dynamodb"
	"github.com/AntonyIS/vision/internal/adapters/http"
	"github.com/AntonyIS/vision/internal/core/services"
)

var env string

func init() {
	config.LoadEnv()
	flag.StringVar(&env, "Environment", "dev", "The environment the application is running")
}

func main() {
	conf := config.AppConfig(env)
	dynamoDB := dynamodb.NewDynamoDB(conf)
	svc := services.NewService(dynamoDB)
	_ = svc
	http.RunServer(conf)
}
