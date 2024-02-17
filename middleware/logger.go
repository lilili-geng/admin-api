package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
)

// 创建logger

func SetupLogger() error {
	var err error
	filepath := viper.GetString("Log.LogFilePath")
	infofilename := viper.GetString("Log.LogInfoFileName")
	warnfilename := viper.GetString("Log.LogWarnFileName")
	fileext := viper.GetString("Log.LogFileExt")
	
	fmt.Println("日志文件路径：", filepath)

	logDir := fmt.Sprintf("%s/%s_%s.%s", filepath, infofilename, time.Now().Format("20060102"), fileext)
	fmt.Println("日志文件路径：", logDir)

	// 检查日志文件夹是否存在，如果不存在则创建
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath, 0755); err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}
	}
	
	// Logger
	Logger, err = GetInitLogger(filepath, infofilename, warnfilename, fileext)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	defer Logger.Sync()
	return nil
}

