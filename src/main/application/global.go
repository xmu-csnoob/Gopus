package application

import (
	"Opus/src/config/mysql"
	"github.com/gin-gonic/gin"
)

var (
	MysqlConfigInstance mysql.Config
	MysqlClientInstance MysqlClient
	StarterInstance     Starter
	HttpEngine          *gin.Engine
)
