package ports

import (
	"context"
	ent "main/internal/auth/domain/entities"
)

// IPostgresqlRepository Auth Domain postgresql interface
type IPostgresqlRepository interface {
	Register(context.Context, ent.RegisterReq) int64
	Login(context.Context, ent.LoginReq) (int64, ent.Auth)
	SaveContact(context.Context, int64, ent.RegisterReq) int64
}
