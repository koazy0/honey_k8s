// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"go.uber.org/zap"
)

type (
	ILogs interface {
		// Cat 新建一个logger，为每个不同的包创建不同的文件夹是，以便后期维护
		Cat(cat string) *zap.SugaredLogger
	}
)

var (
	localLogs ILogs
)

func Logs() ILogs {
	if localLogs == nil {
		panic("implement not found for interface ILogs, forgot register?")
	}
	return localLogs
}

func RegisterLogs(i ILogs) {
	localLogs = i
}
