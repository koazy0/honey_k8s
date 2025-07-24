package model

type LogModel struct {
	Model
	Type        int8   `gorm:"column:type" json:"type"` // 1 登录日志
	IP          string `gorm:"column:ip" json:"ip"`
	Addr        string `gorm:"column:addr" json:"addr"`
	UserID      uint   `gorm:"column:userid" json:"userid"`
	Username    string `gorm:"column:username" json:"username"`
	Pwd         string `gorm:"column:pwd" json:"pwd"`
	LoginStatus bool   `gorm:"column:login_status" json:"login_status"`
	Title       string `gorm:"column:title" json:"title"`
	Level       int8   `gorm:"column:level" json:"level"`
	Content     string `gorm:"column:content" json:"content"`
	ServiceName string `gorm:"column:service_name" json:"service_name"`
}

// 以下结构体用于路由
type CreateLog struct {
}

type CreateLogResponse struct{}

type ListLog struct {
}

type ListLogResponse struct{}

type DeleteLog struct {
}

type DeleteLogResponse struct{}
