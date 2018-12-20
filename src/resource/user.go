package resource

import (
	"config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"model"
	"time"
)

func User(g *gin.RouterGroup){
	g.POST("/token", func(ctx *gin.Context) {
		var loginModel model.LoginModel
		if err := ctx.ShouldBindJSON(&loginModel); err != nil{
			ctx.JSON(400, gin.H{
				"message": "bad request",
				"code": 1,
			})
		}
		if token, err := Verify(loginModel.Username, loginModel.Password); err != nil{
			ctx.JSON(401, gin.H{
				"message": "用户名或密码错误",
				"code": 1,
			})
		}else {
			ctx.JSON(200, gin.H{
				"message": "success",
				"code":    0,
				"token":   token,
				"wish": GetUserByUsername(loginModel.Username).Wish,
			})
		}
	})

}


func Verify(username, password string) (string, error){
	var user model.User
	model.Db.Where("username = ?", username).Find(&user)
	result := CheckPasswordHash(password, user.Password)
	if result{
		j := NewJWT()
		token, err := j.CreateToken(CustomClaims{ID:user.ID, StandardClaims:jwt.StandardClaims{ExpiresAt:time.Now().Add(time.Hour * 1).Unix()}})
		if err != nil{
			config.Error.Println("generateToken error", err.Error())
			return "", err
		}
		return token, err
	}
	return "", errors.New("verify error")
}


func GetUserById(userId uint)model.User{
	var user model.User
	err := model.Db.Where("id=?", userId).Find(&user)
	if err != nil{
		config.Error.Println("GetUserByIde err:", err.Error)
	}
	return user
}


func GetUserByUsername(username string) model.User{
	var user model.User
	err := model.Db.Where("username = ?", username).Find(&user)
	if err != nil{
		config.Error.Println("GetUserByUsername err:", err.Error)
	}
	return user
}