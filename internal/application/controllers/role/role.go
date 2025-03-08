package controllers

import (
	"context"
	"net/http"

	pb "github.com/muharik19/boiler-plate-grpc/api/grpc/api/pb/v1/role"
	"github.com/muharik19/boiler-plate-grpc/configs"
	"github.com/muharik19/boiler-plate-grpc/internal/constant"
	internal "github.com/muharik19/boiler-plate-grpc/internal/pkg/logger"
)

func (s Controllers) CreateRole(ctx context.Context, request *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	ctx = internal.SetIdentifierId(ctx)
	internal.ActivityLogger(ctx, "controllers", "CreateRole", "/api/v1/role", configs.HTTPPost, request, nil)

	role, err := s.RoleService.CreateRole(ctx, request)
	if err != nil {
		response := &pb.CreateRoleResponse{
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
		}

		internal.ErrorLogger(ctx, "controllers", "CreateRole", "/api/v1/role", configs.HTTPPost, request, response, err)

		return response, nil
	}

	internal.ActivityLogger(ctx, "controllers", "CreateRole", "/api/v1/role", configs.HTTPPost, request, role)

	return role, nil
}

func (s Controllers) GetRoleByID(ctx context.Context, request *pb.GetRoleByIDRequest) (*pb.GetRoleByIDResponse, error) {
	ctx = internal.SetIdentifierId(ctx)
	internal.ActivityLogger(ctx, "controllers", "GetRoleByID", "/api/v1/role/{id}", configs.HTTPGet, request, nil)

	role, err := s.RoleService.GetRoleByID(ctx, request)
	if err != nil {
		response := &pb.GetRoleByIDResponse{
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
		}

		internal.ErrorLogger(ctx, "controllers", "GetRoleByID", "/api/v1/role/{id}", configs.HTTPGet, request, response, err)

		return response, nil
	}

	internal.ActivityLogger(ctx, "controllers", "GetRoleByID", "/api/v1/role/{id}", configs.HTTPGet, request, role)

	return role, nil
}

func (s Controllers) GetListRole(ctx context.Context, request *pb.GetListRoleRequest) (*pb.GetListRoleResponse, error) {
	ctx = internal.SetIdentifierId(ctx)
	internal.ActivityLogger(ctx, "controllers", "GetListRole", "/api/v1/role", configs.HTTPGet, request, nil)

	role, err := s.RoleService.GetRoleList(ctx, request)
	if err != nil {
		response := &pb.GetListRoleResponse{
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
		}

		internal.ErrorLogger(ctx, "controllers", "GetListRole", "/api/v1/role", configs.HTTPGet, request, response, err)

		return response, nil
	}

	internal.ActivityLogger(ctx, "controllers", "GetListRole", "/api/v1/role", configs.HTTPGet, request, role)

	return role, nil
}

func (s Controllers) UpdateRole(ctx context.Context, request *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	ctx = internal.SetIdentifierId(ctx)
	internal.ActivityLogger(ctx, "controllers", "UpdateRole", "/api/v1/role/{id}", configs.HTTPPatch, request, nil)

	role, err := s.RoleService.UpdateRole(ctx, request)
	if err != nil {
		response := &pb.UpdateRoleResponse{
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
		}

		internal.ErrorLogger(ctx, "controllers", "UpdateRole", "/api/v1/role/{id}", configs.HTTPPatch, request, response, err)

		return response, nil
	}

	internal.ActivityLogger(ctx, "controllers", "UpdateRole", "/api/v1/role/{id}", configs.HTTPPatch, request, role)

	return role, nil
}

func (s Controllers) DeleteRole(ctx context.Context, request *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	ctx = internal.SetIdentifierId(ctx)
	internal.ActivityLogger(ctx, "controllers", "DeleteRole", "/api/v1/role/{id}", configs.HTTPDel, request, nil)

	role, err := s.RoleService.DeleteRole(ctx, request)
	if err != nil {
		response := &pb.DeleteRoleResponse{
			ResponseCode: constant.FAILED_INTERNAL,
			ResponseDesc: http.StatusText(http.StatusInternalServerError),
		}

		internal.ErrorLogger(ctx, "controllers", "DeleteRole", "/api/v1/role/{id}", configs.HTTPDel, request, response, err)

		return response, nil
	}

	internal.ActivityLogger(ctx, "controllers", "DeleteRole", "/api/v1/role/{id}", configs.HTTPDel, request, role)

	return role, nil
}
