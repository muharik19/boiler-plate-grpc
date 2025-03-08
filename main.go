package main

import (
	"github.com/joho/godotenv"
	cmd "github.com/muharik19/boiler-plate-grpc/cmd/grpc"
	"github.com/muharik19/boiler-plate-grpc/pkg/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Err(err)
	}

	cmd.Init()
}
