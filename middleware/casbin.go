package middleware

import (
	"fmt"
	"os"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/jinzhu/gorm"
)

var (
	Enforcer *casbin.Enforcer
)

// 创建casbin的enforcer
func SetupCasbinEnforcer(db *gorm.DB) {
	if db == nil {
		panic("DB is nil")
	}

	a := gormadapter.NewAdapterByDB(db)

	dir, _ := os.Getwd()
	modelPath := dir + "/config/rbac_model.conf"
	fmt.Println("modelPath:" + modelPath)
	Enforcer = casbin.NewEnforcer(modelPath, a)
	Enforcer.LoadPolicy()
	Enforcer.EnableLog(true)
}
