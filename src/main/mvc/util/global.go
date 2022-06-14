package util

import "Opus/src/config/mysql"

var (
	MysqlConfigInstance        mysql.Config
	MysqlClientInstance        MysqlClient
	ApplicationStarterInstance ApplicationStarter
)
