package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type AdminServicer interface {
	ListAdmins(ctx context.Context) ([]gen.ListAdminsRow, error)
	GetAdmin(adminId int64, ctx context.Context) (gen.GetAdminRow, error)
}

var AdminService AdminServicer = AdminServiceStruct{}

type AdminServiceStruct struct {
}

func (t AdminServiceStruct) ListAdmins(ctx context.Context) ([]gen.ListAdminsRow, error) {
	return db.Queries.ListAdmins(ctx)
}

func (t AdminServiceStruct) GetAdmin(adminId int64, ctx context.Context) (gen.GetAdminRow, error) {
	admin, err := db.Queries.GetAdmin(ctx, adminId)
	if err == sql.ErrNoRows {
		return gen.GetAdminRow{}, UnknownAdminIdError{adminId}
	}
	return admin, nil
}
