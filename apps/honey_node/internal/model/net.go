package model

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net"
)

type NetModel struct {
	Model
	NodeID             uint      `json:"nodeID"`
	NodeModel          NodeModel `gorm:"foreignKey:NodeID" json:"-"`
	Title              string    `gorm:"32" json:"title"`
	Network            string    `gorm:"32" json:"network"` // 网卡
	IP                 string    `gorm:"32" json:"ip"`      // 探针ip
	Mask               int8      `json:"mask"`              // 子网掩码 8-32
	Gateway            string    `gorm:"32" json:"gateway"`
	HostCount          int       `json:"hostCount"`
	HoneyIpCount       int       `json:"honeyIpCount"`
	ScanStatus         int8      `json:"scanStatus"`                    // 扫描状态  0 待扫描  1 扫描完成  2 扫描中
	ScanProgress       float64   `json:"scanProgress"`                  // 扫描进度
	CanUseHoneyIPRange string    `gorm:"256" json:"canUseHoneyIPRange"` // 能够使用的诱捕ip范围
}

func (model *NetModel) Subnet() string {
	return fmt.Sprintf("%s/%d", model.IP, model.Mask)
}

func (model *NetModel) InSubnet(ip string) bool {
	_, _net, _ := net.ParseCIDR(model.Subnet())
	return _net.Contains(net.ParseIP(ip))
}

func (model *NetModel) IpRange() (ipRange []string, err error) {
	//return ip.ParseIPRange(model.CanUseHoneyIPRange)
	//todo 完成utils表
	panic("not implement")
	//return ip.ParseIPRange(model.CanUseHoneyIPRange)
}

func (model *NetModel) BeforeDelete(tx *gorm.DB) error {
	// 是否有诱捕ip
	var count int64
	tx.Model(HoneyIpModel{}).Where("net_id = ?", model.ID).Count(&count)
	if count > 0 {
		return errors.New("存在诱捕ip，不能删除网络")
	}
	// 将启用的网卡，状态归位
	var nodeNet NodeNetworkModel
	err := tx.Take(&nodeNet, "node_id = ? and network = ?", model.NodeID, model.Network).Error
	if err != nil {
		return nil
	}
	var hostList []HostModel
	tx.Find(&hostList, "net_id = ?", model.ID).Delete(&hostList)
	zap.S().Infof("关联删除主机 %d个", len(hostList))
	// 修改状态
	tx.Model(&nodeNet).Update("status", 2)
	return nil
}
