package controller

import (
	"Grpc-crud/app/models"
	//"Grpc-crud/app/pb"
	"Grpc-crud/app/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

func PostBlock(ctx *gin.Context){
	var data models.Binder
	if err := ctx.ShouldBindJSON(&data); err!= nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status":http.StatusBadRequest, "message": "invalid"})
	}
	_, _ = service.WriteRequest(data)
	ctx.JSON(http.StatusOK, gin.H{"status":http.StatusOK, "message": "ok"})
}
func GetBlock(ctx * gin.Context){
	key :=  ctx.Param("partitionkey")
	ctx.JSON(http.StatusOK,gin.H{"status":http.StatusOK, "key": key} )
	
}

func GetBlocks(ctx * gin.Context){
	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
func UpdateBlock(ctx *gin.Context){
	key :=  ctx.Param("partitionkey")
	ctx.JSON(http.StatusOK,gin.H{"status":http.StatusOK, "key": key} )
}

func DeleteBlock(ctx * gin.Context){
	key :=  ctx.Param("partitionkey")
	ctx.JSON(http.StatusOK,gin.H{"status":http.StatusOK, "key": key} )
}