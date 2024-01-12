package main

import (
	"LiadminApi/common"
	"LiadminApi/modules"
	"LiadminApi/routes"
	"LiadminApi/utils/db"
	"fmt"

	_ "LiadminApi/api"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitConfig()
	db.InitDB()

	// 建表
	err := db.DB.AutoMigrate(&modules.UserModule{})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("🚀 Connected Successfully to the table")

	r := gin.Default()

	// 路由
	routes.InitRouter(r)

	common.Init(r, "80", "adminAPI")
}
