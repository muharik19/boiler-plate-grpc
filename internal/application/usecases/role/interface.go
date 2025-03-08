package usecases

import (
	"context"

	pb "github.com/muharik19/boiler-plate-grpc/api/grpc/api/pb/v1/role"
)

type RoleService interface {
	CreateRole(context.Context, *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error)
	GetRoleByID(context.Context, *pb.GetRoleByIDRequest) (*pb.GetRoleByIDResponse, error)
	GetRoleList(context.Context, *pb.GetListRoleRequest) (*pb.GetListRoleResponse, error)
	UpdateRole(context.Context, *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error)
	DeleteRole(context.Context, *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error)
}
