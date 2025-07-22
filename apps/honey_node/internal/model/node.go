package model

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NodeModel struct {
	Model
	Title        string         `gorm:"column:title;size:32" json:"title"` //标题
	Uid          string         `gorm:"column:uid;size:64" json:"uid"`     //uid
	IP           string         `gorm:"column:ip;size:32" json:"ip"`       //IP地址
	Mac          string         `gorm:"column:mac;size:64" json:"mac"`
	Status       int8           `gorm:"column:status" json:"status"`
	NetCount     int            `gorm:"column:net_count" json:"net_count"`
	HoneyIPCount int            `gorm:"column:honey_ip_count" json:"honey_ip_count"`
	Resource     NodeResource   `gorm:"column:resource;serializer:json" json:"resource"` //列名为Resource，序列化成json进行
	SystemInfo   NodeSystemInfo `gorm:"column:system_info;serializer:json" json:"system_info"`
}

type NodeResource struct {
	CpuCount              int     `json:"cpu_count"`    //cpu数
	CpuUseRate            float64 `json:"cpu_use_rate"` //cpu使用率
	MemTotal              float64 `json:"mem_total"`    //mem总量，单位GB
	MemUseRate            float64 `json:"mem_use_rate"` //mem使用率
	DiskTotal             int64   `json:"disk_total"`
	DiskUseRate           float64 `json:"disk_use_rate"`
	NodePath              string  `json:"node_path"`               // 节点的部署目录
	NodeResourceOccupancy int64   `json:"node_resource_occupancy"` // 节点目录的资源占用
}

type NodeSystemInfo struct {
	HostName            string `json:"host_name"`
	DistributionVersion string `json:"distribution_version"` // 发行版本
	CoreVersion         string `json:"core_version"`         // 内核版本
	SystemType          string `json:"system_type"`          // 系统类型
	StartTime           string `json:"start_time"`           // 启动时间
	NodeVersion         string `json:"node_version"`
	NodeCommit          string `json:"node_commit"`
}

func (n *NodeModel) BeforeDelete(tx *gorm.DB) error {
	// 诱捕转发
	var list []HoneyPortModel
	err := tx.Find(&list, "node_id = ?", n.ID).Delete(&list).Error
	if err != nil {
		return err
	}
	zap.S().Infof("关联诱捕转发 %d", len(list))

	// 诱捕ip
	var ipList []HoneyIpModel
	err = tx.Find(&list, "node_id = ?", n.ID).Delete(&ipList).Error
	if err != nil {
		return err
	}
	zap.S().Infof("关联诱捕ip %d", len(ipList))

	//节点网络
	var netList []NetModel
	err = tx.Find(&list, "node_id = ?", n.ID).Delete(&netList).Error
	if err != nil {
		return err
	}
	zap.S().Infof("关联网络 %d", len(netList))

	// 节点网卡
	var networkList []NodeNetworkModel
	err = tx.Find(&list, "node_id = ?", n.ID).Delete(&networkList).Error
	if err != nil {
		return err
	}
	zap.S().Infof("关联节点网卡 %d", len(networkList))
	// 节点
	return nil
}

type NodeNetworkModel struct {
	Model
	NodeID    uint      `gorm:"column:node_id;" json:"node_id"`
	NodeModel NodeModel `gorm:"foreignKey:node_id" json:"-"`
	Network   string    `gorm:"column:network;32" json:"network"`
	IP        string    `gorm:"column:ip;32" json:"ip"`   // 探针ip
	Mask      int8      `gorm:"column:mask;" json:"mask"` // 子网掩码 8-32
	Gateway   string    `gorm:"column:gateway;32" json:"gateway"`
	Status    int8      `gorm:"column:status;" json:"status"` // 是否启用 1 启用 2 未启用
}

func (n *NodeNetworkModel) BeforeDelete(tx *gorm.DB) error {
	// 先找有没有网络
	if n.Status == 2 {
		// 未启用
		return nil
	}
	var net NetModel
	err := tx.Take(&net, "node_id = ? and network = ?", n.NodeID, n.Network).Error
	if err != nil {
		// 未启用
		return nil
	}

	// 判断有没有诱捕ip
	var count int64
	tx.Model(HoneyIpModel{}).Where("net_id = ?", net.ID).Count(&count)
	if count > 0 {
		return errors.New("此网卡的网络存在诱捕ip，不可删除")
	}

	// 关联删除网络表和主机表
	var hostList []HostModel
	tx.Find(&hostList, "net_id = ?", net.ID).Delete(&hostList)
	tx.Delete(&net)
	zap.S().Infof("关联删除主机记录 %d", len(hostList))
	zap.S().Infof("关联删除网络 %s", net.Title)
	return nil
}
