package util

import "log"

type ApplicationStarter struct {
}

func (applicationStarter *ApplicationStarter) InitDatabase() {
	log.Printf("\033[1;34;40m%s\033[0m\n\n", "INFO : 数据库初始化")
	MysqlClientInstance.Initialize()
}
func (applicationStarter *ApplicationStarter) InitMiddleware() {
	log.Printf("\033[1;34;40m%s\033[0m\n\n", "INFO : 中间件初始化")
}
func (applicationStarter *ApplicationStarter) InitHttpInterfaces() {
	log.Printf("\033[1;34;40m%s\033[0m\n\n", "INFO : HTTP请求初始化")
}
func (applicationStarter *ApplicationStarter) Initialize() {
	log.Printf("\033[1;34;40m%s\033[0m\n\n", "INFO : 服务启动中")
	ApplicationStarterInstance.InitDatabase()
	ApplicationStarterInstance.InitMiddleware()
	ApplicationStarterInstance.InitHttpInterfaces()
	log.Printf("\033[1;34;40m%s\033[0m\n\n", "INFO : 服务启动成功")
}
