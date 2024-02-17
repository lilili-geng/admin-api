package db

import (
	"LiadminApi/middleware"
	"LiadminApi/modules"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
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
func InitDB() error {

	// newlogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
	// 	SlowThreshold: time.Second, // sql
	// 	LogLevel:      logger.Info, // 级别
	// 	Colorful:      true,        // 颜色
	// })

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

	DB, err = gorm.Open("mysql", db)
	if err != nil {
		return err
	}

	// 全局禁用表名复数
	DB.SingularTable(true)
	//打开sql日志
	DB.LogMode(true)
	fmt.Println("🚀 Connected Successfully to the Database")
	return nil
}

func Loginit() {
	InitConfig()

	err := middleware.SetupLogger()

	if err != nil {
		log.Fatalf("failed to set up logger: %v", err)
	}

	err = InitDB()

	if err != nil {
		middleware.Logger.Fatalf("db :: %v", err)
		log.Fatalf("failed to initialize DB: %v", err)
	}

	// 初始化数据库表结构
	DB.AutoMigrate(&modules.SysUserModule{}, &modules.SysRoleModule{}, &modules.SysUserRole{})


	err = middleware.SetupAccessLogger()

	if err != nil {
		log.Fatalf("failed to set up access logger: %v", err)
	}

	middleware.SetupCasbinEnforcer(DB)

	middleware.Logger.Infof("------应用init结束")
}
