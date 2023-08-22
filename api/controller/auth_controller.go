package controller

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/june21x/project-a-api/util"
// )

// type LoginReq struct {
// 	// PlayerName string `json:"playerName" binding:"required"`
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

// func Login(ctx *gin.Context) {
// 	var reqBodyJson LoginReq

// 	if err := ctx.ShouldBindJSON(&reqBodyJson); err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, util.NewErrorResponse(err))
// 		return
// 	}

// 	// TODO login
// 	// player, err := authservice.Login( reqBodyJson.Email, reqBodyJson.Password)
// 	// if err != nil {
// 	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(err))
// 	// 	return
// 	// }

// 	// ctx.JSON(http.StatusOK, player)
// }

// func Logout(ctx *gin.Context) {

// 	// TODO logout
// 	// err := authservice.Logout(token)
// 	// if err != nil {
// 	// 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, util.NewErrorResponse(err))
// 	// 	return
// 	// }

// 	// ctx.JSON(http.StatusOK)
// }
