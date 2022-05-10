package model

type Blog_article_tag struct {
	Id         int    `gorm:"id" json:"id"`
	ArticleId  int    `gorm:"article_id" json:"article_id"`   // 文章 ID
	TagId      int    `gorm:"tag_id" json:"tag_id"`           // 标签 ID
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	CreatedBy  string `gorm:"created_by" json:"created_by"`   // 创建人
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	ModifiedBy string `gorm:"modified_by" json:"modified_by"` // 修改人
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}

type Blog_tag struct {
	Id         int    `gorm:"id" json:"id"`
	Name       string `gorm:"name" json:"name"`               // 标签名称
	CreatedOn  int    `gorm:"created_on" json:"created_on"`   // 创建时间
	CreatedBy  string `gorm:"created_by" json:"created_by"`   // 创建人
	ModifiedOn int    `gorm:"modified_on" json:"modified_on"` // 修改时间
	ModifiedBy string `gorm:"modified_by" json:"modified_by"` // 修改人
	DeletedOn  int    `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      int    `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
	State      int    `gorm:"state" json:"state"`             // 状态 0 为禁用、1 为启用
}

type Blog_article struct {
	Id            int    `gorm:"id" json:"id"`
	Title         string `gorm:"title" json:"title"`                     // 文章标题
	Desc          string `gorm:"desc" json:"desc"`                       // 文章简述
	CoverImageUrl string `gorm:"cover_image_url" json:"cover_image_url"` // 封面图片地址
	Content       string `gorm:"content" json:"content"`                 // 文章内容
	CreatedOn     int    `gorm:"created_on" json:"created_on"`           // 创建时间
	CreatedBy     string `gorm:"created_by" json:"created_by"`           // 创建人
	ModifiedOn    int    `gorm:"modified_on" json:"modified_on"`         // 修改时间
	ModifiedBy    string `gorm:"modified_by" json:"modified_by"`         // 修改人
	DeletedOn     int    `gorm:"deleted_on" json:"deleted_on"`           // 删除时间
	IsDel         int    `gorm:"is_del" json:"is_del"`                   // 是否删除 0 为未删除、1 为已删除
	State         int    `gorm:"state" json:"state"`                     // 状态 0 为禁用、1 为启用
}
