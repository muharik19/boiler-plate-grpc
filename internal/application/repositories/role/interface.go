package repositories

import (
	"context"

	role "github.com/muharik19/boiler-plate-grpc/internal/domain/entities/role"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/utils"
)

type RoleRepository interface {
	CreateRole(context.Context, role.Role) (*role.Role, error)
	GetRoleByID(context.Context, string) (*role.Role, error)
	GetRoleExistsByName(context.Context, string) (*role.Role, *bool, error)
	GetRoleListWithPagination(context.Context, utils.PaginationRequest, map[string]string) (*[]role.Role, *int, error)
	PatchRole(context.Context, role.Role) error
	DeleteRole(context.Context, role.Role) error
}
