package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"main/config"
	"main/pkg/utils/env"
	"time"
)

type TokenClaim struct {
	AuthId      int64  `json:"auth_id"`
	Lang        string `json:"lang"`
	ParentId    int64  `json:"parent_id,omitempty"`
	ManagerId   int64  `json:"manager_id,omitempty"`
	AccountType int8   `json:"acc_type,omitempty"`
	UserType    int8   `json:"user_type,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	UserTitle   string `json:"user_title"`
	UserName    string `json:"user_name"`
	IsDemo      bool   `json:"is_demo,omitempty"`
	UniqueId    string `json:"unique_id"`
	Status      int8   `json:"status,omitempty"`
	Exp         int64  `json:"exp"`
	Iat         int64  `json:"iat"`
	jwt.RegisteredClaims
}

func GenerateToken(cfg *config.Config, claims TokenClaim) (string, error) {
	signingKey := []byte(cfg.Server.APP_SECRET)
	claims.Iat = time.Now().Unix()
	claims.Exp = time.Now().Local().Add(time.Hour * time.Duration(env.EnvInt("JWT_TOKEN_EXPIRE_TIME"))).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
