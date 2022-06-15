package controller

import (
	"Opus/src/main/mvc/service"
	"Opus/src/main/mvc/util"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
	engine       *gin.Engine
	imageService service.ImageService
}

func (imageController *ImageController) Start(engine *gin.Engine) {
	imageController.GetHttpEngine(engine)
	imageController.RegisterHttpInterfaces()
}
func (imageController *ImageController) GetHttpEngine(engine *gin.Engine) {
	util.LoggerInstance.PrintInfo("ImageController:获取Gin Engine")
	imageController.engine = engine
}
func (imageController *ImageController) RegisterHttpInterfaces() {
	util.LoggerInstance.PrintInfo("ImageController:注册Http接口")
	imageController.engine.GET("/test",
		func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello gopus",
			})
		})
}
