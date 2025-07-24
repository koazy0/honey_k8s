package model

type HostModel struct {
	Model
	NodeID    uint      `gorm:"column:node_id;" json:"node_id"`
	NodeModel NodeModel `gorm:"foreignKey:node_id" json:"-"`
	NetID     uint      `gorm:"column:net_id;" json:"net_id"`
	NetModel  NetModel  `gorm:"foreignKey:net_id" json:"-"`
	IP        string    `gorm:"column:ip;32" json:"ip"`
	Mac       string    `gorm:"column:mac;64" json:"mac"`
	Manuf     string    `gorm:"column:manuf;64" json:"manuf"` // 厂商信息
}
type HostTemplateModel struct {
	Model
	Title    string               `gorm:"column:title;size:32" json:"title"`
	PortList HostTemplatePortList `gorm:"column:port_list;serializer:json" json:"port_list"`
}
type HostTemplatePortList []HostTemplatePort
type HostTemplatePort struct {
	Port      int  `json:"port"`
	ServiceID uint `json:"serviceID"`
}

// 以下结构体用于路由

type CreateHost struct {
}

type CreateHostResponse struct{}

type ListHost struct {
}

type ListHostResponse struct{}

type DeleteHost struct {
}

type DeleteHostResponse struct{}
