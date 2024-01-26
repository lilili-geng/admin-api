package main

import (
	"LiadminApi/middleware"
	"LiadminApi/modules"
	"LiadminApi/routes"
	"LiadminApi/utils"
	"LiadminApi/utils/db"
	"fmt"

	_ "LiadminApi/api"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitConfig()
	db.InitDB()

	// 建表
	err := db.DB.AutoMigrate(&modules.SysUserModule{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("🚀 Connected Successfully to the table")

	r := gin.Default()

	// 中间件
	r.Use(middleware.Cors())

	// 路由
	routes.InitRouter(r)

	utils.Init(r, "80", "adminAPI")
}
