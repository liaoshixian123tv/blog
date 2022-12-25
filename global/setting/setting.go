package setting

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
	"fmt"
	"log"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)

var (
	//	yaml Section
	Server   = "Server"
	App      = "App"
	Database = "Database"
)

func init() {
	ServerSetting = new(setting.ServerSettings)
	AppSetting = new(setting.AppSettings)
	DatabaseSetting = new(setting.DatabaseSettings)
	Logger = new(logger.Logger)

	setupSetting()
	setupLogger()
}

// setup初始化設定
func setupSetting() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}
	settings.ReadSection(Server, &ServerSetting)
	if err != nil {
		return err
	}
	settings.ReadSection(Database, &DatabaseSetting)
	if err != nil {
		return err
	}
	settings.ReadSection(App, &AppSetting)
	if err != nil {
		return err
	}

	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	fmt.Printf("ServerSetting: %+v \n", ServerSetting)
	fmt.Printf("AppSetting: %+v \n", AppSetting)
	fmt.Printf("Database: %+v \n", DatabaseSetting)
	return nil
}

func setupLogger() {
	fileName := AppSetting.LogSavePath + "/" + AppSetting.LogFileName + AppSetting.LogFileExt
	fmt.Println(fileName)
	Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

}
