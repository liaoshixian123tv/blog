package setting

import (
	"blog/pkg/setting"
	"fmt"
	"time"
)

var (
	ServerSetting   *setting.ServerSettings
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
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
	setup()
}

// setup初始化設定
func setup() error {
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
