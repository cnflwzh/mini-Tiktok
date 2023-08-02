package jwt

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/dgrijalva/jwt-go"
)

const (
	// 密钥，用于签名Token
	secretKey = "miniTiktok"
	// Token过期时间，240小时
	expirationTime = 240 * time.Hour
	//expirationTime = time.Second
)

type CustomClaims struct {
	Username string `json:"username"`
	UserID   int    `json:"userId"`
	jwt.StandardClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(username string, userID int) (string, error) {
	hlog.Info("Begin to generate Token for userID:", userID, " username:", username)
	// 创建一个自定义的Claims
	claims := &CustomClaims{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	// 使用HMAC SHA256算法签名Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名Token并获取完整的Token字符串
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		hlog.Error("Generate Token Fail ID:", userID, " username:", username)
		return "", err
	}
	hlog.Info("Token Generated:", signedToken)
	return signedToken, nil
}

// ParseToken 解析JWT Token，如果成功返回userID，如果失败返回错误
func ParseToken(tokenString string) (int, error) {
	// 解析Token，同时验证签名和过期时间
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	// 检查Token中的Claims是否是我们自定义的CustomClaims类型
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("invalid token")
}

// JWTAuthMiddleware 由token生成用户ID的中间件
func JWTAuthMiddleware() []app.HandlerFunc {
	return []app.HandlerFunc{
		func(ctx context.Context, c *app.RequestContext) {
			// 从请求头中获取Token
			tokenString, ok := c.Get("token")
			if tokenString == nil {
				tokenString = string(c.FormValue("token"))
				if tokenString != nil {
					ok = true
				} else {
					ok = false
				}
			}
			if !ok {
				hlog.Error("Get Token from header fail")
				c.AbortWithStatus(401)
				return
			}
			// 解析Token
			userID, err := ParseToken(tokenString.(string))
			if err != nil {
				hlog.Error("Parse Token Fail:", err)
				c.AbortWithStatus(401)
				return
			}
			// 将用户ID存入上下文
			c.Set("userID", userID)
			c.Next(ctx)
		}}
}

// JWTGenMiddleware 由用户ID生成token的中间件
func JWTGenMiddleware() []app.HandlerFunc {
	return []app.HandlerFunc{
		func(ctx context.Context, c *app.RequestContext) {
			// 从上下文中获取用户ID
			userID, ok := c.Get("userID")
			if !ok {
				hlog.Error("Get userID from context fail")
				c.AbortWithStatus(401)
				return
			}
			// 从上下文中获取用户名
			username, ok := c.Get("username")
			if !ok {
				hlog.Error("Get username from context fail")
				c.AbortWithStatus(401)
				return
			}
			// 生成Token
			token, err := GenerateToken(username.(string), userID.(int))
			if err != nil {
				hlog.Error("Generate token fail:", err)
				c.AbortWithStatus(401)
				return
			}
			// 将Token存入上下文
			c.Set("token", token)
			c.Next(ctx)
		}}
}
