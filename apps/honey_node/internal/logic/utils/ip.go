package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"scaffold/internal/service"
	"strconv"
	"strings"
)

type (
	sIP struct{}
)

var insIP = sIP{}

func IP() *sIP {
	return &insIP
}

func init() {

	service.Logs().Info("Init utils.ip success!")
	//service.RegisterIP(IP())
}

// HasLocalIPAddr 判断ip是否是本地ip
func (s *sIP) HasLocalIPAddr(_ip string) bool {
	ip := net.ParseIP(_ip)
	if ip.IsPrivate() {
		return true
	}
	if ip.IsLoopback() {
		return true
	}
	return false
}

// ParseIPRange 解析IP范围字符串，支持单个IP和IP段，返回IP字符串列表
func (s *sIP) ParseIPRange(ipRange string) ([]string, error) {
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

// IncrementIP 将IP地址加1
func (s *sIP) IncrementIP(ip net.IP) net.IP {
	if ip == nil {
		return nil
	}

	// 复制IP地址，避免修改原IP
	newIP := make(net.IP, len(ip))
	copy(newIP, ip)

	// 处理IPv4地址
	if ip4 := newIP.To4(); ip4 != nil {
		for i := 3; i >= 0; i-- {
			newIP[i]++
			if newIP[i] > 0 {
				break
			}
		}
		return newIP
	}

	// 处理IPv6地址
	for i := len(newIP) - 1; i >= 0; i-- {
		newIP[i]++
		if newIP[i] > 0 {
			break
		}
	}
	return newIP
}

// DecrementIP 将IP地址减1
func (s *sIP) DecrementIP(ip net.IP) net.IP {
	if ip == nil {
		return nil
	}

	newIP := make(net.IP, len(ip))
	copy(newIP, ip)

	if ip4 := newIP.To4(); ip4 != nil {
		for i := 3; i >= 0; i-- {
			newIP[i]--
			if newIP[i] < 255 {
				break
			}
		}
		return newIP
	}

	for i := len(newIP) - 1; i >= 0; i-- {
		newIP[i]--
		if newIP[i] < 255 {
			break
		}
	}
	return newIP
}

// BroadcastIP 计算CIDR块的广播地址
func (s *sIP) BroadcastIP(network *net.IPNet) net.IP {
	ip := network.IP.To4()
	if ip == nil {
		// 处理IPv6广播地址 (实际上IPv6没有广播地址)
		return nil
	}

	mask := network.Mask
	result := make(net.IP, len(ip))

	for i := 0; i < len(ip); i++ {
		result[i] = ip[i] | ^mask[i]
	}

	return result
}

// FormatIPRange 格式化IP范围为字符串
func (s *sIP) FormatIPRange(start, end net.IP) string {
	return fmt.Sprintf("%s-%s", start, end)
}

// intToIP 将整数转换为IPv4地址
// e.g.:192.168.2.1=>
func intToIP(ipInt uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		(ipInt>>24)&0xff,
		(ipInt>>16)&0xff,
		(ipInt>>8)&0xff,
		ipInt&0xff)
}

// iPToInt 将IPv4地址转换为整数
func (s *sIP) iPToInt(ipStr string) (ipInt uint32) {
	ip := net.ParseIP(ipStr).To4()
	if ip == nil {
		service.Logs().Errorf("invalid IPv4 address: %q", ipStr)
	}
	// 按大端（网络字节序）直接转 uint32
	// 因为这时存到IP这个切片里就是大端序，直接转成uint32就行了
	return binary.BigEndian.Uint32(ip)

}

func (s *sIP) ParseCIDRGetUseIPRange(cidr string) (r string, err error) {
	ipObj, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		err = errors.New("无效的网段")
		return
	}
	mask, _ := ipNet.Mask.Size()
	// 转换为IPv4地址
	ip4 := ipObj.To4()
	if ip4 == nil {
		err = errors.New("不是有效的IPv4地址")
		return
	}

	// 处理掩码小于24的情况，取第一个C段
	if mask < 24 {
		ipParts := strings.Split(ip4.String(), ".")
		if len(ipParts) != 4 {
			err = errors.New("无效的IPv4地址格式")
			return
		}
		// 构建第一个C段地址
		firstC := fmt.Sprintf("%s.%s.%s.0", ipParts[0], ipParts[1], ipParts[2])
		ip4 = net.ParseIP(firstC).To4()
		mask = 24
	}

	// 计算网络地址和广播地址
	ipInt := s.iPToInt(string(ip4))
	maskBits := uint(32 - mask)
	networkInt := ipInt & (^uint32(0) << maskBits)
	broadcastInt := networkInt | (^uint32(0) >> (32 - maskBits))

	// 计算可用IP范围
	firstUsable := networkInt + 1
	lastUsable := broadcastInt - 1

	// 输出结果
	r = fmt.Sprintf("%s-%s", intToIP(firstUsable), intToIP(lastUsable))
	return
}
