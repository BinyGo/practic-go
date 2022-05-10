package model

import "github.com/practic-go/gin/blog/pkg/app"

type Tag struct {
	*Model
	Name  string `gorm:"name" json:"name"`   // 标签名称
	State int    `gorm:"state" json:"state"` // 状态 0 为禁用、1 为启用
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
