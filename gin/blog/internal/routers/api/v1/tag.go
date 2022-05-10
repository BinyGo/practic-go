package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/practic-go/gin/blog/global"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}
func (t Tag) List(c *gin.Context) {
	fmt.Println(global.ServerSetting)
	fmt.Println(global.AppSetting)
	fmt.Println(global.DatabaseSetting)
}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
