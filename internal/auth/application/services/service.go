package services

import (
	"context"
	"main/config"
	ent "main/internal/auth/domain/entities"
	"main/internal/auth/domain/ports"
	"main/pkg/logger"
)

var (
	err error
)

const (
	collection = "examples"
)

// serviceAuth Auth Service
type serviceAuth struct {
	cfg    *config.Config
	pgRepo ports.IPostgresqlRepository
	logger logger.Logger
}

// NewAuthService Auth domain service constructor
func NewAuthService(cfg *config.Config, pgRepo ports.IPostgresqlRepository, logger logger.Logger) ports.IService {
	return &serviceAuth{cfg: cfg, pgRepo: pgRepo, logger: logger}
}

func (t serviceAuth) Register(ctx context.Context, req_body ent.RegisterReq) (record int64) {

	record = t.pgRepo.Register(ctx, req_body)

	return
}

func (t serviceAuth) Login(ctx context.Context, req_body ent.LoginReq) (record int64, data ent.Auth) {

	record, data = t.pgRepo.Login(ctx, req_body)

	return
}
