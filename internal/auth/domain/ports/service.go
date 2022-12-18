package ports

import (
	"context"
	ent "main/internal/auth/domain/entities"
)

// IService Auth domain service interface
type IService interface {
	Register(context.Context, ent.RegisterReq) int64
	Login(context.Context, ent.LoginReq) (int64, ent.Auth)
}
