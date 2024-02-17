package main

import (
	"LiadminApi/middleware"
	"LiadminApi/routes"
	"LiadminApi/utils"
	"LiadminApi/utils/db"
	"fmt"

	_ "LiadminApi/api"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	db.Loginit()

	fmt.Println("Logger in main:", middleware.Logger) // Add this line

	middleware.Logger.Infof("------应用main函数开始")

	gin.SetMode(viper.GetString("Log.RunMode"))

	r := gin.Default()


	routes.InitRouter(r)

	utils.Init(r, "80", "adminAPI")
}
