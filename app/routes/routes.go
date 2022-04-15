package routes


import (
	"net/http"
	"Grpc-crud/app/controller"
	"github.com/gin-gonic/gin"
)

func Url_maps() {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Is Alive....!!")
	})
	v1 := r.Group("/v1")
	{
		v1.POST("/postblock", controller.PostBlock)
		v1.GET("/getblock/:partitionkey", controller.GetBlock)
		v1.GET("/getblocks", controller.GetBlocks)
		v1.DELETE("/deleteblock/:partitionkey", controller.DeleteBlock)
		v1.PATCH("/updateblock/:partitionkey", controller.UpdateBlock)
	}
}
