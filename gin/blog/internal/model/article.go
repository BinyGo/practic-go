package model

type Article struct {
	*Model
	Title         string `gorm:"title" json:"title"`                     // 文章标题
	Desc          string `gorm:"desc" json:"desc"`                       // 文章简述
	CoverImageUrl string `gorm:"cover_image_url" json:"cover_image_url"` // 封面图片地址
	Content       string `gorm:"content" json:"content"`                 // 文章内容
	State         int    `gorm:"state" json:"state"`                     // 状态 0 为禁用、1 为启用
}

func (t Article) TableName() string {
	return "blog_article"
}
