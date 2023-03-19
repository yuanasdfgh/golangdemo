package main

import (
	"server/initialize"
)

func main() {
	// 初始化配置 / viper
	//global.Viper = initialize.InitializeConfig()
	//
	////初始化日志
	//global.Log = initialize.InitializeLog() // 初始化zap日志库
	//zap.ReplaceGlobals(global.Log)
	// 初始化数据库
	//global.DB = initialize.InitializeDB()

	//global.Log.Info("log init success!")
	// 初始化验证器
	// 初始化Redis

	// 初始化文件系统

	// 初始化计划任务

	// 启动服务器
	initialize.Router()
}
