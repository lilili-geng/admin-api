package main

import (
	"LiadminApi/modules"
	"LiadminApi/utils/db"
	"fmt"

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

	r.Run()
	fmt.Println("启动成功")
}
