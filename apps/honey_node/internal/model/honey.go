package model

//蜜罐IP的结构

type HoneyIpModel struct {
	Model
	NodeID    uint      `json:"nodeID"`
	NodeModel NodeModel `gorm:"foreignKey:NodeID" json:"-"`
	NetID     uint      `json:"netID"`
	NetModel  NetModel  `gorm:"foreignKey:NetID" json:"-"`
	IP        string    `gorm:"32" json:"ip"`
	Mac       string    `gorm:"64" json:"mac"`
	Network   string    `gorm:"32" json:"network"` // 网卡
	Status    int8      `json:"status"`            // 1 创建中 2 运行中 3 失败  4 删除中
	ErrorMsg  string    `json:"errorMsg"`
}

type HoneyPortModel struct {
	Model
	NodeID       uint         `json:"nodeID"`
	NodeModel    NodeModel    `gorm:"foreignKey:NodeID" json:"-"`
	NetID        uint         `json:"netID"`
	NetModel     NetModel     `gorm:"foreignKey:NetID" json:"-"`
	HoneyIpID    uint         `json:"honeyIpID"`
	HoneyIpModel HoneyIpModel `gorm:"foreignKey:HoneyIpID" json:"-"`
	ServiceID    uint         `json:"serviceID"` // 服务id
	ServiceModel ServiceModel `gorm:"foreignKey:ServiceID" json:"-"`
	Port         int          `json:"port"`                 // 服务的端口
	DstIP        string       `gorm:"size:32" json:"dstIP"` // 目标ip
	DstPort      int          `json:"dstPort"`              // 目标端口
	Status       int8         `json:"status"`
}
