package resource

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

func KingAngAngle(group *gin.RouterGroup){
	group.GET("/king", func(ctx *gin.Context) {
		currentUser := ctx.MustGet("current_user").(model.User)
		king := getKingOrAngle("k", currentUser)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"code": 0,
			"data": gin.H{
				"king": king.Name,
				"king_username": king.Username,
				"king_wish": king.Wish,
			},
		})
	})

	group.GET("/angle", func(ctx *gin.Context) {
		currentUser := ctx.MustGet("current_user").(model.User)
		angle := getKingOrAngle("a", currentUser)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"code": 0,
			"data": gin.H{
				"angle_name": angle.Name,
				"angle_blessing": angle.Blessing,
			},
		})
	})
}

func getKingOrAngle(t string, user model.User)(model.User){
	if t == "k"{
		return GetUserByUsername(user.KingUsername)
	}else {
		var myAngle model.User
		model.Db.Where("king_username = ?", user.Username).Find(&myAngle)
		return myAngle
	}
}