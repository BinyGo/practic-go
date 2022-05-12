package global

import (
	"github.com/practic-go/gin/blog/pkg/logger"
	"github.com/practic-go/gin/blog/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)

var (
	DBEngine   *gorm.DB
	Logger     *logger.Logger
	JWTSetting *setting.JWTSettingS
)
