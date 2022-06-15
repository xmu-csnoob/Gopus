package database

import (
	"Opus/src/main/mvc/util"
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type MysqlClient struct {
	GormDBClient *gorm.DB
}

func (mysqlClient *MysqlClient) Initialize() {
	mysqlClient.ReadConfigs()
	mysqlClient.ConnectToDatabase()
}
func (mysqlClient *MysqlClient) ReadConfigs() {
	iniPath := "src/config/mysql/mysql.ini"
	util.LoggerInstance.PrintInfo("加载数据库配置文件")
	cfg, err := ini.Load(iniPath)
	if err != nil {
		fmt.Println(err)
	}
	err = cfg.Section("mysql").MapTo(&MysqlConfigInstance)
}

func (mysqlClient *MysqlClient) ConnectToDatabase() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		MysqlConfigInstance.UserName, MysqlConfigInstance.Password, MysqlConfigInstance.Host, MysqlConfigInstance.Port,
		MysqlConfigInstance.Database, MysqlConfigInstance.Charset, MysqlConfigInstance.ParseTime, MysqlConfigInstance.Loc)
	// 连接额外配置信息
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   MysqlConfigInstance.TablePre, //表前缀
			SingularTable: true,                         //使用单数表名，启用该选项时，`User` 的表名应该是 `user`而不是users
		},
	}
	util.LoggerInstance.PrintInfo("连接数据库")
	// 打印SQL设置
	if MysqlConfigInstance.PrintSqlLog {
		slowSqlTime, err := time.ParseDuration(MysqlConfigInstance.SlowSqlTime)
		if err != nil {
			fmt.Println(err)
		}
		loggerNew := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: slowSqlTime, //慢SQL阈值
			LogLevel:      logger.Info,
			Colorful:      true, // 彩色打印开启
		})
		gormConfig.Logger = loggerNew
		// 建立连接
		mysqlClient.GormDBClient, err = gorm.Open(mysql.Open(dataSourceName), gormConfig)
		if err != nil {
			panic(err)
		}
		// 设置连接池信息
		db, err2 := mysqlClient.GormDBClient.DB()
		if err != nil {
			panic(err2)
		}
		// 设置空闲连接池中连接的最大数量
		db.SetMaxIdleConns(MysqlConfigInstance.MaxIdleConn)
		// 设置打开数据库连接的最大数量
		db.SetMaxOpenConns(MysqlConfigInstance.MaxOpenConn)
		// 设置了连接可复用的最大时间
		duration, err := time.ParseDuration(MysqlConfigInstance.MaxLifeTime)
		if err != nil {
			panic(err)
		}
		db.SetConnMaxLifetime(duration)
	}
}
