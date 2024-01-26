package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// 测试读取配置文件
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("init config")
}

// 连接数据库 && 并且根据module建立对应的表
func InitDB() {
	newlogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, // sql
		LogLevel:      logger.Info, // 级别
		Colorful:      true,        // 颜色
	})

	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	database := viper.GetString("mysql.database")
	charset := viper.GetString("mysql.charset")
	parseTime := viper.GetBool("mysql.parseTime")
	loc := viper.GetString("mysql.loc")

	db := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, parseTime, loc)

	var err error

	DB, err = gorm.Open(mysql.Open(db), &gorm.Config{Logger: newlogger})

	if err != nil {
		panic(err)
	}

	fmt.Println("🚀 Connected Successfully to the Database")
	
}
