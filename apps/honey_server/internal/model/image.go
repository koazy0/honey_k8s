package model

type ImageModel struct {
	Model
	ImageName     string `gorm:"column:image_name" json:"image_name"`
	Title         string `gorm:"column:title" json:"title"`
	Port          int    `gorm:"column:port" json:"port"`
	DockerImageID string `gorm:"column:docker_image_id;size:32" json:"docker_image_id"`
	Tag           string `gorm:"column:tag" json:"tag"`
	Agreement     int8   `gorm:"column:agreement" json:"agreement"`
	ImagePath     string `gorm:"column:image_path" json:"image_path"` // 镜像文件
	Status        int8   `gorm:"column:status" json:"status"`
	Logo          string `gorm:"column:logo" json:"logo"` // 镜像的logo
	Desc          string `gorm:"column:desc" json:"desc"` // 镜像描述
}
