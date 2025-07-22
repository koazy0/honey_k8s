package model

type MatrixTemplateModel struct {
	Model
	Title            string           `gorm:"size:32" json:"title"`
	HostTemplateList HostTemplateList `gorm:"serializer:json" json:"hostTemplateList"` // 主机模板列表
}
type HostTemplateList []HostTemplateInfo
type HostTemplateInfo struct {
	HostTemplateID uint `json:"hostTemplateID"`
	Weight         int  `json:"weight"`
}
