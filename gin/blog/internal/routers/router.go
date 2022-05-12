package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/practic-go/gin/blog/docs"
	"github.com/practic-go/gin/blog/global"
	"github.com/practic-go/gin/blog/internal/middleware"
	"github.com/practic-go/gin/blog/internal/routers/api"
	v1 "github.com/practic-go/gin/blog/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()
	//curl -X POST http://127.0.0.1:8999/upload/file -F "file=@/golang/src/github.com/practic-go/README.md" -F type=1
	//curl -X POST http://127.0.0.1:8999/upload/file -F "file=@/golang/src/github.com/practic-go/gin/blog/storage/uploads/biny.jpg" -F type=1
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.POST("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
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
