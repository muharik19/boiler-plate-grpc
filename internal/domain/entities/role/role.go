package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	ID        string       `json:"id" bun:"type:varchar(30),pk,notnull"`
	RoleName  string       `json:"roleName" bun:"type:varchar(50),notnull"`
	CreatedAt time.Time    `json:"createdAt" bun:",nullzero,notnull,default:current_timestamp"`
	CreatedBy string       `json:"createdBy" bun:"type:varchar(30),notnull"`
	UpdatedAt time.Time    `json:"updatedAt" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedBy string       `json:"updatedBy" bun:",nullzero"`
	DeletedAt bun.NullTime `json:"deletedAt" bun:",soft_delete,nullzero"`
	DeletedBy string       `json:"deletedBy" bun:"type:varchar(30),nullzero"`
}
