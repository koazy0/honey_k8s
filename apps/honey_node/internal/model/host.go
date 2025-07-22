package model

type HostModel struct {
	Model
	NodeID    uint      `json:"nodeID"`
	NodeModel NodeModel `gorm:"foreignKey:NodeID" json:"-"`
	NetID     uint      `json:"netID"`
	NetModel  NetModel  `gorm:"foreignKey:NetID" json:"-"`
	IP        string    `gorm:"32" json:"ip"`
	Mac       string    `gorm:"64" json:"mac"`
	Manuf     string    `gorm:"64" json:"manuf"` // 厂商信息
}
type HostTemplateModel struct {
	Model
	Title    string               `gorm:"size:32" json:"title"`
	PortList HostTemplatePortList `gorm:"serializer:json" json:"portList"`
}
type HostTemplatePortList []HostTemplatePort
type HostTemplatePort struct {
	Port      int  `json:"port"`
	ServiceID uint `json:"serviceID"`
}
