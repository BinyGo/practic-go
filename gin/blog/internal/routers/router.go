package routers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/practic-go/gin/blog/docs"
	"github.com/practic-go/gin/blog/global"
	"github.com/practic-go/gin/blog/internal/middleware"
	"github.com/practic-go/gin/blog/internal/routers/api"
	v1 "github.com/practic-go/gin/blog/internal/routers/api/v1"
	"github.com/practic-go/gin/blog/pkg/limiter"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "api",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	// 本地环境
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger()) //控制台输出访问日志
		r.Use(gin.Recovery())
		r.Use(middleware.Tracing())
		r.Use(middleware.AccessLog())
		r.Use(middleware.RateLimiter(methodLimiters))
		r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	} else {
		r.Use(gin.Recovery())
		r.Use(middleware.Tracing())
		r.Use(middleware.AccessLog())
		r.Use(middleware.RateLimiter(methodLimiters))
		r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	}

	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()

	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	//apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}

	return r
}
