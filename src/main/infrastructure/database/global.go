package database

import (
	"Opus/src/config/mysql"
)

var (
	MysqlConfigInstance mysql.Config
	MysqlClientInstance MysqlClient
)
