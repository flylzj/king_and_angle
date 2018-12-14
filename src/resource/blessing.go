package resource

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

func Blessing(g *gin.RouterGroup){
	g.GET("/king", func(ctx *gin.Context) {
		currentUser := ctx.MustGet("current_user").(model.User)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"code": 0,
			"data": gin.H{
				"blessing": currentUser.Blessing,
			},
		})
	})

	g.POST("/king", func(ctx *gin.Context) {
		var blessing model.Blessing
		currentUser := ctx.MustGet("current_user").(model.User)
		if err := ctx.ShouldBindJSON(&blessing); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "bad requests",
				"code": 1,
			})
		}else {
			if blessing.Blessing == ""{
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "blessing to short",
					"code": 1,
				})
			}else{
				currentUser.Blessing = blessing.Blessing
				model.Db.Save(&currentUser)
				ctx.JSON(http.StatusOK, gin.H{
					"message": "success",
					"code": 0,
				})
			}
		}
	})

	g.GET("/angle", func(ctx *gin.Context) {
		currentUser := ctx.MustGet("current_user").(model.User)

		angle := getKingOrAngle("a", currentUser)
		if angle.ID == 0{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "no angle",
				"code": 1,
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
			"code": 0,
			"data": gin.H{
				"blessing": angle.Blessing,
			},
		})

	})
}
