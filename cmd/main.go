package main

import (
	"fmt"

	"github.com/AntonyIS/vision1.0/config"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("Vision Portal")
}
