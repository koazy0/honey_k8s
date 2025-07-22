package model

import (
	"bytes"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"honey_node/internal/common"
	"net"
	"strconv"
	"strings"
)

type NetModel struct {
	Model
	NodeID             uint      `gorm:"column:node_id" json:"node_id"`
	NodeModel          NodeModel `gorm:"foreignKey:node_id" json:"-"`
	Title              string    `gorm:"column:title;32" json:"title"`
	Network            string    `gorm:"column:network;32" json:"network"` // 网卡
	IP                 string    `gorm:"column:ip;32" json:"ip"`           // 探针ip
	Mask               int8      `gorm:"column:mask" json:"mask"`          // 子网掩码 8-32
	Gateway            string    `gorm:"column:gateway;32" json:"gateway"`
	HostCount          int       `gorm:"column:host_count" json:"host_count"`
	HoneyIpCount       int       `gorm:"column:honey_ip_count" json:"honey_ip_count"`
	ScanStatus         int8      `gorm:"column:scan_status" json:"scan_status"`                           // 扫描状态  0 待扫描  1 扫描完成  2 扫描中
	ScanProgress       float64   `gorm:"column:scan_progress" json:"scan_progress"`                       // 扫描进度
	CanUseHoneyIPRange string    `gorm:"column:can_use_honey_ip_range;256" json:"can_use_honey_ip_range"` // 能够使用的诱捕ip范围
}

func (model *NetModel) Subnet() string {
	return fmt.Sprintf("%s/%d", model.IP, model.Mask)
}

func (model *NetModel) InSubnet(ip string) bool {
	_, _net, _ := net.ParseCIDR(model.Subnet())
	return _net.Contains(net.ParseIP(ip))
}

func (model *NetModel) IpRange() (ipRange []string, err error) {
	return parseIPRange(model.CanUseHoneyIPRange)
}

var logger = common.Logs().Cat("model")

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
	logger.Infof("关联删除主机 %d个", len(hostList))
	// 修改状态
	tx.Model(&nodeNet).Update("status", 2)
	return nil
}

// ParseIPRange 解析IP范围字符串，支持单个IP和IP段，返回IP字符串列表，放出来避免循环调用
func parseIPRange(ipRange string) ([]string, error) {
	var result []string
	// 按逗号分割不同的IP或IP段
	segments := strings.Split(ipRange, ",")

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)
		if segment == "" {
			continue
		}

		// 检查是否包含连字符（IP段）
		if strings.Contains(segment, "-") {
			parts := strings.SplitN(segment, "-", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("无效的IP段格式: %s", segment)
			}

			startIPStr := strings.TrimSpace(parts[0])
			endPart := strings.TrimSpace(parts[1])

			startIP := net.ParseIP(startIPStr)
			if startIP == nil {
				return nil, fmt.Errorf("无效的起始IP: %s", startIPStr)
			}

			// 处理IPv4地址段
			if ipv4 := startIP.To4(); ipv4 != nil {
				startIP = ipv4
				// 尝试将结束部分解析为完整IP或数字
				var endIP net.IP
				if endIP = net.ParseIP(endPart); endIP != nil {
					endIP = endIP.To4()
					if endIP == nil {
						return nil, fmt.Errorf("无效的结束IP: %s", endPart)
					}
				} else {
					// 尝试将结束部分解析为数字（IP的最后一个八位组）
					endNum, err := strconv.Atoi(endPart)
					if err != nil || endNum < 0 || endNum > 255 {
						return nil, fmt.Errorf("无效的结束部分: %s", endPart)
					}
					endIP = make(net.IP, len(startIP))
					copy(endIP, startIP)
					endIP[len(endIP)-1] = byte(endNum)
				}

				// 生成IP范围
				for cmp := bytes.Compare(startIP, endIP); cmp <= 0; cmp = bytes.Compare(startIP, endIP) {
					result = append(result, startIP.String())
					// 递增IP
					for i := len(startIP) - 1; i >= 0; i-- {
						startIP[i]++
						if startIP[i] > 0 {
							break
						}
					}
				}
			} else {
				return nil, fmt.Errorf("IPv6范围解析暂不支持")
			}
		} else {
			// 处理单个IP
			ip := net.ParseIP(segment)
			if ip == nil {
				return nil, fmt.Errorf("无效的IP地址: %s", segment)
			}
			result = append(result, ip.String())
		}
	}

	return result, nil
}
