package model

type ServiceModel struct {
	Model
	Title        string     `gorm:"column:title" json:"title"`
	Agreement    int8       `gorm:"column:agreement" json:"agreement"`
	ImageID      uint       `gorm:"column:image_id" json:"image_id"`
	ImageModel   ImageModel `gorm:"foreignKey:image_id" json:"-"`
	IP           string     `gorm:"column:ip" json:"ip"`
	Port         int        `gorm:"column:port" json:"port"`
	Status       int8       `gorm:"column:status" json:"status"`
	HoneyIPCount int        `gorm:"column:honey_ip_count" json:"honey_ip_count"`
	ContainerID  string     `gorm:"column:container_id" json:"container_id"` // 容器id
}
