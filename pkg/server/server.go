package server

import (
	"context"
	"main/config"
	"main/pkg/logger"
)

const (
	certFile        = "ssl/server.crt"
	keyFile         = "ssl/server.pem"
	maxHeaderBytes  = 1 << 20
	gzipLevel       = 5
	stackSize       = 1 << 10 // 1 KB
	csrfTokenHeader = "X-CSRF-Token"
	bodyLimit       = "2M"
)

// server
type server struct {
	cfg    *config.Config
	ctx    *context.Context
	logger logger.Logger
}

// NewServer constructor
func NewServer(cfg *config.Config, ctx *context.Context, logger logger.Logger) *server {
	return &server{cfg: cfg, ctx: ctx, logger: logger}
}
