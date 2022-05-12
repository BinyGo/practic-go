package global

import (
	"github.com/opentracing/opentracing-go"
	"github.com/practic-go/gin/blog/pkg/logger"
	"github.com/practic-go/gin/blog/pkg/setting"
	"gorm.io/gorm"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
)

var (
	DBEngine   *gorm.DB
	Logger     *logger.Logger
	Tracer     opentracing.Tracer
	TracerSpan opentracing.Span
)
