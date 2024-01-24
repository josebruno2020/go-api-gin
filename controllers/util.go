package controllers

import "github.com/gin-gonic/gin"

func sendJsonError(ctx *gin.Context, err HttpError) {
	ctx.JSON(err.Code, err)
}
