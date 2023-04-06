package main

import (
	"flag"
	"fmt"

	"github.com/AntonyIS/vision1.0/config"
	"github.com/AntonyIS/vision1.0/internal/adapters/database/dynamodb"
	"github.com/AntonyIS/vision1.0/internal/adapters/http"
	"github.com/AntonyIS/vision1.0/internal/core/services"
)

var env string

func init() {
	config.LoadEnv()
	flag.StringVar(&env, "Environment", "dev", "The environment the application is running")
}

func setup() {
	conf := config.AppConfig(env)
	dynamoDB := dynamodb.NewDynamoDB(conf)
	svc := services.NewService(dynamoDB)
	_ = svc

	http.RunServer(conf)

}

func main() {
	fmt.Println("Vision Portal")
	setup()
}
