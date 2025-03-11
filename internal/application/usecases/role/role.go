package usecases

import (
	"context"
	"math"
	"net/http"
	"time"

	pb "github.com/muharik19/boiler-plate-grpc/api/grpc/api/pb/v1/role"
	repositories "github.com/muharik19/boiler-plate-grpc/internal/application/repositories/role"
	"github.com/muharik19/boiler-plate-grpc/internal/constant"
	entities "github.com/muharik19/boiler-plate-grpc/internal/domain/entities/role"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/utils"
	"github.com/uptrace/bun/schema"
)

type roleService struct {
	RoleRepository repositories.RoleRepository
}

func NewRoleService() *roleService {
	return &roleService{}
}

func (s *roleService) SetRoleRepository(roleRepository repositories.RoleRepository) *roleService {
	s.RoleRepository = roleRepository
	return s
}

func (s *roleService) CreateRole(ctx context.Context, request *pb.CreateRoleRequest) (*pb.CreateRoleResponse, error) {
	body := request.GetBody()

	id, err := utils.GeneratePK("ROL")
	if err != nil {
		return nil, err
	}

	_, exists, err := s.RoleRepository.GetRoleExistsByName(ctx, body.GetName())
	if err != nil {
		return nil, err
	}

	if *exists {
		return &pb.CreateRoleResponse{
			ResponseCode: constant.FAILED_EXIST,
			ResponseDesc: "Name already exist",
		}, nil
	}

	payload := entities.Role{
		ID:        *id,
		RoleName:  body.GetName(),
		CreatedBy: "system",
	}

	role, err := s.RoleRepository.CreateRole(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRoleResponse{
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.CreateRoleResponse_ResponseData{
			Id: role.ID,
		},
	}, nil
}

func (s *roleService) GetRoleByID(ctx context.Context, request *pb.GetRoleByIDRequest) (*pb.GetRoleByIDResponse, error) {
	role, err := s.RoleRepository.GetRoleByID(ctx, request.GetId())
	if err != nil {
		return &pb.GetRoleByIDResponse{
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
		}, nil
	}

	return &pb.GetRoleByIDResponse{
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.GetRoleByIDResponse_ResponseData{
			Id:   role.ID,
			Name: role.RoleName,
		},
	}, nil
}

func (s *roleService) GetRoleList(ctx context.Context, request *pb.GetListRoleRequest) (*pb.GetListRoleResponse, error) {
	pagination := utils.Pagination(int(request.GetLimit()), int(request.GetPage()), request.GetField(), request.GetSort())

	items, count, err := s.RoleRepository.GetRoleListWithPagination(ctx, pagination, request.GetFilter())
	if err != nil {
		return &pb.GetListRoleResponse{
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
		}, nil
	}

	if *count == 0 {
		return &pb.GetListRoleResponse{
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
		}, nil
	}

	var roles []*pb.GetListRoleResponse_ResponseData_Role
	for _, item := range *items {
		role := &pb.GetListRoleResponse_ResponseData_Role{
			Id:        item.ID,
			Name:      item.RoleName,
			CreatedAt: item.CreatedAt.Format("02-01-2006 15:04:05"),
			CreatedBy: item.CreatedBy,
			UpdatedAt: item.UpdatedAt.Format("02-01-2006 15:04:05"),
			UpdatedBy: item.UpdatedBy,
		}
		roles = append(roles, role)
	}

	return &pb.GetListRoleResponse{
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.GetListRoleResponse_ResponseData{
			Page:      int32(pagination.Page),
			Limit:     int32(pagination.Limit),
			Total:     int32(*count),
			TotalPage: int32(math.Ceil(float64(*count) / float64(pagination.Limit))),
			Roles:     roles,
		},
	}, nil
}

func (s *roleService) UpdateRole(ctx context.Context, request *pb.UpdateRoleRequest) (*pb.UpdateRoleResponse, error) {
	id := request.GetId()
	body := request.GetBody()

	_, err := s.RoleRepository.GetRoleByID(ctx, id)
	if err != nil {
		return &pb.UpdateRoleResponse{
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
		}, nil
	}

	role, exists, err := s.RoleRepository.GetRoleExistsByName(ctx, body.GetName())
	if err != nil {
		return nil, err
	}

	if *exists {
		if role.ID != id {
			return &pb.UpdateRoleResponse{
				ResponseCode: constant.FAILED_EXIST,
				ResponseDesc: "Name already exist",
			}, nil
		}
	}

	payload := entities.Role{
		ID:        id,
		RoleName:  body.GetName(),
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}

	err = s.RoleRepository.PatchRole(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateRoleResponse{
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.UpdateRoleResponse_ResponseData{
			Id:   id,
			Name: body.GetName(),
		},
	}, nil
}

func (s *roleService) DeleteRole(ctx context.Context, request *pb.DeleteRoleRequest) (*pb.DeleteRoleResponse, error) {
	id := request.GetId()

	role, err := s.RoleRepository.GetRoleByID(ctx, id)
	if err != nil {
		return &pb.DeleteRoleResponse{
			ResponseCode: constant.FAILED_NOT_FOUND,
			ResponseDesc: http.StatusText(http.StatusNotFound),
		}, nil
	}

	payload := entities.Role{
		ID: id,
		DeletedAt: schema.NullTime{
			Time: time.Now(),
		},
		DeletedBy: "system",
	}

	err = s.RoleRepository.DeleteRole(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteRoleResponse{
		ResponseCode: constant.SUCCESS,
		ResponseDesc: http.StatusText(http.StatusOK),
		ResponseData: &pb.DeleteRoleResponse_ResponseData{
			Id:   id,
			Name: role.RoleName,
		},
	}, nil
}
