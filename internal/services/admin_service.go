package services

import (
	db "aurora/database"
	gen "aurora/database/gen"
	"database/sql"

	"context"
)

type AdminServiceStruct struct {
}

var AdminService AdminServiceStruct

func (t AdminServiceStruct) ListAdmins(ctx context.Context) ([]gen.ListAdminsRow, error) {
	return db.Queries.ListAdmins(ctx)
}

func (t AdminServiceStruct) GetAdmin(AdminId int64, ctx context.Context) (gen.GetAdminRow, error) {
	Admin, err := db.Queries.GetAdmin(ctx, AdminId)
	if err == sql.ErrNoRows {
		return gen.GetAdminRow{}, &UnknownAdminIdError{AdminId}
	}
	return Admin, nil
}