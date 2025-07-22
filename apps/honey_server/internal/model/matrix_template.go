package model

type MatrixTemplateModel struct {
	Model
	Title            string           `gorm:"column:title;size:32" json:"title"`
	HostTemplateList HostTemplateList `gorm:"column:host_template_list;serializer:json" json:"host_template_list"` // 主机模板列表
}
type HostTemplateList []HostTemplateInfo
type HostTemplateInfo struct {
	HostTemplateID uint `json:"host_template_id"`
	Weight         int  `json:"weight"`
}
