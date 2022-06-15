package application

import (
	"Opus/src/main/mvc/controller"
	"Opus/src/main/mvc/util"
	"github.com/gin-gonic/gin"
)

type Starter struct {
	ImageControllerInstance controller.ImageController
}

func (applicationStarter *Starter) InitDatabase() {
	util.LoggerInstance.PrintInfo("数据库初始化")
	MysqlClientInstance.Initialize()
}
func (applicationStarter *Starter) InitMiddleware() {
	util.LoggerInstance.PrintInfo("中间件初始化")
}
func (applicationStarter *Starter) InitHttpInterfaces() {
	util.LoggerInstance.PrintInfo("HTTP接口初始化")
	HttpEngine = gin.Default()
	applicationStarter.ImageControllerInstance.Start(HttpEngine)
	HttpEngine.Run()
}
func (applicationStarter *Starter) Initialize() {
	util.LoggerInstance.PrintInfo("服务启动中")
	StarterInstance.InitDatabase()
	StarterInstance.InitMiddleware()
	StarterInstance.InitHttpInterfaces()
	util.LoggerInstance.PrintInfo("服务启动成功")
}
