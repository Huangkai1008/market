package jwtauth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"market/internal/pkg/config"
)

// Options jwt配置可选项
type Options struct {
	*config.Jwt
}

// Claims 签名
type Claims struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

// JwtAuth jwt认证
type JwtAuth struct {
	opts *Options
}

// New 创建认证实例
func New(opts *Options) *JwtAuth {
	return &JwtAuth{opts: opts}
}

// GenerateToken 生成令牌
func (a *JwtAuth) GenerateToken(userID uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(a.opts.JwtExpireDuration * time.Second)

	claims := Claims{
		UserId:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  nowTime.Unix(),
			ExpiresAt: expireTime.Unix(),
			Issuer:    a.opts.JwtIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(a.opts.JwtIssuer))
	return token, err
}

// ParseToken 解析令牌
func (a *JwtAuth) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.opts.JwtIssuer), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// ParseUserID 从上下文中获取UserID
func (a *JwtAuth) ParseUserID(ctx *gin.Context) (uint, error) {
	if userId, exist := ctx.Get("userId"); exist {
		return userId.(uint), nil
	} else {
		return 0, errors.New("授权中间件异常")
	}
}
