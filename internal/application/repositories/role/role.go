package repositories

import (
	"context"
	"fmt"

	role "github.com/muharik19/boiler-plate-grpc/internal/domain/entities/role"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/database"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/logger"
	"github.com/muharik19/boiler-plate-grpc/internal/pkg/utils"
	"github.com/uptrace/bun"
)

type roleRepository struct {
	db *bun.DB
}

func NewRoleRepository() *roleRepository {
	return &roleRepository{
		db: database.DbBun,
	}
}

func (r roleRepository) CreateRole(ctx context.Context, role role.Role) (*role.Role, error) {
	logger.ActivityLogger(ctx, "repositories", "CreateRole", "", "", nil, nil)

	_, err := r.db.NewInsert().Model(&role).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r roleRepository) GetRoleByID(ctx context.Context, id string) (*role.Role, error) {
	logger.ActivityLogger(ctx, "repositories", "GetRoleByID", "", "", nil, nil)

	role := new(role.Role)
	err := r.db.NewSelect().Model(role).Column("id", "role_name").Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r roleRepository) GetRoleExistsByName(ctx context.Context, name string) (*role.Role, *bool, error) {
	logger.ActivityLogger(ctx, "repositories", "GetRoleExistsByName", "", "", nil, nil)

	role := new(role.Role)
	exists, err := r.db.NewSelect().Model(role).Column("id", "role_name").Where("role_name = ?", name).Exists(ctx)
	if err != nil {
		return nil, nil, err
	}

	return role, &exists, nil
}

func (r roleRepository) GetRoleListWithPagination(ctx context.Context, pagination utils.PaginationRequest, where map[string]string) (*[]role.Role, *int, error) {
	logger.ActivityLogger(ctx, "repositories", "GetRoleListWithPagination", "", "", nil, nil)

	var field, sort string
	var roles []role.Role

	query := r.db.NewSelect().Model(&roles).Column("id", "role_name", "created_at", "created_by", "updated_at", "updated_by")

	if where["id"] != "" {
		query = query.Where("id = ?", where["id"])
	}

	if where["name"] != "" {
		name := fmt.Sprintf("%%%s%%", where["name"])
		query = query.Where("role_name ILIKE ?", name)
	}

	if pagination.Field != "" {
		if pagination.Field == "id" {
			field = "INITCAP(id)"
		} else if pagination.Field == "name" {
			field = `INITCAP("role_name")`
		} else {
			field = "created_at"
		}
	} else {
		field = "created_at"
	}

	if pagination.Sort != "" {
		sort = pagination.Sort
	} else {
		sort = "DESC"
	}

	offset := (pagination.Page - 1) * pagination.Limit
	orderBy := fmt.Sprintf("%s %s", field, sort)
	count, err := query.Order(orderBy).Limit(pagination.Limit).Offset(offset).ScanAndCount(ctx)
	if err != nil {
		return nil, nil, err
	}
	return &roles, &count, nil
}

func (r roleRepository) PatchRole(ctx context.Context, role role.Role) error {
	logger.ActivityLogger(ctx, "repositories", "PatchRole", "", "", nil, nil)

	_, err := r.db.NewUpdate().Model(&role).Column("role_name", "updated_at", "updated_by").WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r roleRepository) DeleteRole(ctx context.Context, role role.Role) error {
	logger.ActivityLogger(ctx, "repositories", "DeleteRole", "", "", nil, nil)

	_, err := r.db.NewUpdate().Model(&role).Column("deleted_at", "deleted_by").WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
