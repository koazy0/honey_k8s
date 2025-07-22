package model

//蜜罐IP的结构，用于描述蜜罐的状态

// HoneyIpModel 蜜罐的IP的结构体
type HoneyIpModel struct {
	Model
	NodeID    uint      `gorm:"column:node_id" json:"node_id"`
	NodeModel NodeModel `gorm:"foreignKey:node_id" json:"-"`
	NetID     uint      `gorm:"column:net_id" json:"net_id"`
	NetModel  NetModel  `gorm:"foreignKey:net_id" json:"-"`
	IP        string    `gorm:"32;column:ip" json:"ip"`
	Mac       string    `gorm:"64;column:mac" json:"mac"`
	Network   string    `gorm:"32;column:network" json:"network"` // 网卡
	Status    int8      `gorm:"column:status" json:"status"`      // 1 创建中 2 运行中 3 失败  4 删除中
	ErrorMsg  string    `gorm:"column:error_msg" json:"error_msg"`
}

// HoneyPortModel 蜜罐端口的结构体
type HoneyPortModel struct {
	Model
	NodeID       uint         `gorm:"column:node_id" json:"node_id"`
	NodeModel    NodeModel    `gorm:"foreignKey:node_id" json:"-"`
	NetID        uint         `gorm:"column:net_id" json:"net_id"`
	NetModel     NetModel     `gorm:"foreignKey:net_id" json:"-"`
	HoneyIpID    uint         `gorm:"column:honey_ip_id" json:"honey_ip_id"`
	HoneyIpModel HoneyIpModel `gorm:"foreignKey:honey_ip_id" json:"-"`
	ServiceID    uint         `gorm:"column:service_id" json:"service_id"` // 服务id
	ServiceModel ServiceModel `gorm:"foreignKey:service_id" json:"-"`
	Port         int          `gorm:"column:port" json:"port"`             // 服务的端口
	DstIP        string       `gorm:"column:dst_ip;size:32" json:"dst_ip"` // 目标ip
	DstPort      int          `gorm:"column:dst_port" json:"dst_port"`     // 目标端口
	Status       int8         `gorm:"column:status" json:"status"`
}
