package resource

import (
	"config"
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
				"king_wish_status": king.WishFinished,
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
				"angle_username": angle.Username,
			},
		})
	})

	group.POST("/wish", func(ctx *gin.Context) {
		var wishModel model.WishModel
		if err := ctx.ShouldBindJSON(&wishModel); err != nil{
			ctx.JSON(400, gin.H{
				"message": "bad request",
				"code": 1,
			})
		}
		if wishModel.Wish == ""{
			ctx.JSON(400, gin.H{
				"message": "bad request",
				"code": 1,
			})
		}
		currentUser := ctx.MustGet("current_user").(model.User)
		currentUser.Wish = wishModel.Wish
		err := model.Db.Save(&currentUser)
		if err.Error != nil{
			ctx.JSON(200, gin.H{
				"message": "save wish failed",
				"code": 1,
			})
		}
		ctx.JSON(200, gin.H{
			"message": "success",
			"code": 0,
		})
	})

	group.GET("/wish", func(ctx *gin.Context) {
		currentUser := ctx.MustGet("current_user").(model.User)
		ctx.JSON(200, gin.H{
			"message": "success",
			"code": 0,
			"wish": currentUser.Wish,
			"wish_status": currentUser.WishFinished,
		})
	})

	group.POST("/wish_status", func(ctx *gin.Context) {
		currentUser := ctx.MustGet("current_user").(model.User)
		var wishStatus model.WishStatusModel
		if err := ctx.ShouldBindJSON(&wishStatus); err != nil{
			ctx.JSON(400, gin.H{
				"message": "bad request" + err.Error(),
				"code": "1",
			})
		}
		status := [2]uint{0, 1}
        config.Info.Println(wishStatus.WishStatus == status[1])
		if wishStatus.WishStatus != status[0] && wishStatus.WishStatus != status[1]{
			ctx.JSON(400, gin.H{
				"message": "bad request",
				"code": "1",
			})
		}else{
			currentUser.WishFinished = wishStatus.WishStatus
			if err := model.Db.Save(&currentUser).Error; err != nil{
				ctx.JSON(200, gin.H{
					"message": err.Error(),
					"code": "1",
				})
			}else{
				ctx.JSON(200, gin.H{
					"message": "success",
					"code": "0",
				})
			}
		}
	})
}

func getKingOrAngle(t string, user model.User)(model.User){
	if t == "k"{
		return GetUserByUsername(user.KingUsername)
	}else {
		var myAngle model.User
		err := model.Db.Where("king_username = ?", user.Username).Find(&myAngle)
		if err.Error != nil{
			config.Error.Println("GetKingOrAngle(angle) err:", err.Error)
		}
		return myAngle
	}
}