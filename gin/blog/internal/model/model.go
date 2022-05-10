package model

type Model struct {
	Id         uint32 `gorm:"id" json:"id"`
	CreatedBy  string `gorm:"created_by" json:"created_by"`   // 创建人
	ModifiedBy string `gorm:"modified_by" json:"modified_by"` // 修改人
	CreatedOn  uint32 `gorm:"created_on" json:"created_on"`   // 创建时间
	ModifiedOn uint32 `gorm:"modified_on" json:"modified_on"` // 修改时间
	DeletedOn  uint32 `gorm:"deleted_on" json:"deleted_on"`   // 删除时间
	IsDel      uint8  `gorm:"is_del" json:"is_del"`           // 是否删除 0 为未删除、1 为已删除
}
