package jwt

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
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
	UserID   int64  `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(username string, userID int64) (string, error) {
	hlog.Info("Begin to generate Token for userID:", userID, " username:", username)
	// 创建一个自定义的Claims
	claims := CustomClaims{
		username,
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
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
func ParseToken(tokenString string) (int64, error) {
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
			// 从请求中获取Token
			token := c.QueryArgs().Peek("token")
			if token == nil {
				hlog.Error("Get Token Fail")
				c.AbortWithStatus(401)
				return
			}
			// 将Token转换为字符串
			tokenString := string(token)
			hlog.Info("Get Token:", tokenString)
			// 解析Token
			userID, err := ParseToken(tokenString)
			if err != nil {
				hlog.Error("Parse Token Fail:", err)
				c.AbortWithStatus(401)
				return
			}
			if c.QueryArgs().Peek("user_id") != nil {
				c.Request.SetQueryString(fmt.Sprintf("token_user_id=%d&", userID) + string(c.Request.QueryString()))
			} else {
				// 将用户ID存入params
				c.Request.SetQueryString(fmt.Sprintf("user_id=%d&", userID) + string(c.Request.QueryString()))
			}
			// 看一下params
			hlog.Info("Params:", string(c.Request.QueryString()))
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
			token, err := GenerateToken(username.(string), userID.(int64))
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

// CheckLoginMiddleware 用于鉴权用户是否登录
func CheckLoginMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从请求中获取Token
		token := c.QueryArgs().Peek("token")
		if token == nil {
			hlog.Error("Token not provided")
			c.AbortWithStatus(401)
			return
		}
		// 将Token转换为字符串
		tokenString := string(token)
		// 解析Token
		userID, err := ParseToken(tokenString)
		if err != nil {
			hlog.Error("Invalid Token:", err)
			c.AbortWithStatus(401)
			return
		}
		// 将用户ID存入上下文，以便后续处理使用
		c.Set("loginUserId", userID)
		c.Next(ctx)
	}
}
