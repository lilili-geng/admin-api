package middleware

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	AccessLogger *zap.SugaredLogger
)

// 创建Accesslogger
func SetupAccessLogger() error {
	var err error
	filepath := viper.GetString("AccessLog.LogFilePath")
	filename := viper.GetString("AccessLog.LogFileName")
	fileext := viper.GetString("AccessLog.LogFileExt")

	AccessLogger, err = GetInitAccessLogger(filepath, filename, fileext)

	if err != nil {
		return err
	}
	defer AccessLogger.Sync()
	return nil
}
