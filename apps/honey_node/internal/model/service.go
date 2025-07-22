package model

type ServiceModel struct {
	Model
	Title        string     `json:"title"`
	Agreement    int8       `json:"agreement"`
	ImageID      uint       `json:"imageID"`
	ImageModel   ImageModel `gorm:"foreignKey:ImageID" json:"-"`
	IP           string     `json:"ip"`
	Port         int        `json:"port"`
	Status       int8       `json:"status"`
	HoneyIPCount int        `json:"honeyIPCount"`
	ContainerID  string     `json:"containerID"` // 容器id
}
