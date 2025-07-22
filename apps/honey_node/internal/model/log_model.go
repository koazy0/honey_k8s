package model

type LogModel struct {
	Model
	Type        int8   `json:"type"` // 1 登录日志
	IP          string `json:"ip"`
	Addr        string `json:"addr"`
	UserID      uint   `json:"userID"`
	Username    string `json:"username"`
	Pwd         string `json:"pwd"`
	LoginStatus bool   `json:"loginStatus"`
	Title       string `json:"title"`
	Level       int8   `json:"level"`
	Content     string `json:"content"`
	ServiceName string `json:"serviceName"`
}
