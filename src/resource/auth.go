package resource

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/pbkdf2"
	"net/http"
	"strings"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"message":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": 1,
					"msg":    "授权已过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("current_user", GetUserById(claims.ID))
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "christmas"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    uint `json:"user_id"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}


func CheckToken(token string)(*CustomClaims, bool){
	j := NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == TokenExpired {
			return nil, false
		}
		return nil, false
	}
	return claims, true
}

var (
	iterations  = 1000
)

func HashInternal(salt string, password string) string {
	hash := pbkdf2.Key([]byte(password), []byte(salt), iterations, 20, sha1.New)
	return hex.EncodeToString(hash)
}

func CheckPasswordHash(password string, hash string) bool {
	if strings.Count(hash, "$") < 2 {
		return false
	}
	pwdHashLists := strings.Split(hash, "$")
	return pwdHashLists[2] == HashInternal(pwdHashLists[1], password)
}
