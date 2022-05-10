package model

type ArticleTag struct {
	*Model
	ArticleId int `gorm:"article_id" json:"article_id"` // 文章 ID
	TagId     int `gorm:"tag_id" json:"tag_id"`         // 标签 ID
}

func (t ArticleTag) TableName() string {
	return "blog_article_tag"
}
