package api

import (
	"log"
	"net"

	pb "github.com/muharik19/boiler-plate-grpc/api/grpc/api/pb/v1/role"
	controllersRole "github.com/muharik19/boiler-plate-grpc/internal/application/controllers/role"
	repositoriesRole "github.com/muharik19/boiler-plate-grpc/internal/application/repositories/role"
	usecasesRole "github.com/muharik19/boiler-plate-grpc/internal/application/usecases/role"
	"github.com/muharik19/boiler-plate-grpc/pkg/logger"
	global "github.com/muharik19/boiler-plate-grpc/pkg/utils"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

func NewGrpcServer() error {
	logger.Configure()
	port, err := net.Listen("tcp", ":"+*global.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// REPOSITORY
	roleRepository := repositoriesRole.NewRoleRepository()

	// USECASE
	roleUseCase := usecasesRole.NewRoleService().SetRoleRepository(roleRepository)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery())),
		grpc.StreamInterceptor(apmgrpc.NewStreamServerInterceptor()),
	)

	controllers := controllersRole.CreateControllers(s, roleUseCase)

	pb.RegisterRoleServer(s, controllers)
	// register grpc service server

	logger.Infof("Serving gRPC on 0.0.0.0: %v", *global.Getenv("GRPC_PORT"))
	s.Serve(port)

	return nil
}
