// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"net"
)

type (
	IExample interface {
		Example()
	}
	IIP interface {
		// HasLocalIPAddr 判断ip是否是本地ip
		HasLocalIPAddr(_ip string) bool
		// ParseIPRange 解析IP范围字符串，支持单个IP和IP段，返回IP字符串列表
		ParseIPRange(ipRange string) ([]string, error)
		// IncrementIP 将IP地址加1
		IncrementIP(ip net.IP) net.IP
		// DecrementIP 将IP地址减1
		DecrementIP(ip net.IP) net.IP
		// BroadcastIP 计算CIDR块的广播地址
		BroadcastIP(network *net.IPNet) net.IP
		// FormatIPRange 格式化IP范围为字符串
		FormatIPRange(start net.IP, end net.IP) string
		ParseCIDRGetUseIPRange(cidr string) (r string, err error)
	}
	IJwt interface {
		// GenerateToken 生成一个带有 userID的 JWT，默认有效期 12 小时
		GenerateToken(ctx context.Context, userID string) (string, error)
		ParseToken(tokenString string) (UserID string, err error)
		// HashPassword 用 SHA-256(salt + password) 方式加密
		HashPassword(password string, salt string) (passwordEncrypt string)
		// GenerateSalt 生成随机盐值，长度固定为16个字节
		GenerateSalt() (string, error)
	}
	IMigrations interface {
		Migrate(ctx context.Context)
	}
)

var (
	localExample    IExample
	localIP         IIP
	localJwt        IJwt
	localMigrations IMigrations
)

func Example() IExample {
	if localExample == nil {
		panic("implement not found for interface IExample, forgot register?")
	}
	return localExample
}

func RegisterExample(i IExample) {
	localExample = i
}

func IP() IIP {
	if localIP == nil {
		panic("implement not found for interface IIP, forgot register?")
	}
	return localIP
}

func RegisterIP(i IIP) {
	localIP = i
}

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}

func Migrations() IMigrations {
	if localMigrations == nil {
		panic("implement not found for interface IMigrations, forgot register?")
	}
	return localMigrations
}

func RegisterMigrations(i IMigrations) {
	localMigrations = i
}
