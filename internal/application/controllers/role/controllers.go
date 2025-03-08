package controllers

import (
	usecasesRole "github.com/muharik19/boiler-plate-grpc/internal/application/usecases/role"
	"google.golang.org/grpc"
)

type Controllers struct {
	Grpc        *grpc.Server
	RoleService usecasesRole.RoleService
}

func CreateControllers(
	grpc *grpc.Server,
	role usecasesRole.RoleService,
) *Controllers {
	return &Controllers{
		grpc,
		role,
	}
}
