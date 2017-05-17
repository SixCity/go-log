package main

import "time"

type RecordLogs struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	// 记录类型
	Type string `gorm:"not null" form:"type" query:"type" json:"type"`
	// 客户端栈
	Model string `gorm:"not null" form:"model" query:"model" json:"model"`
	// 客户端记录内容
	Content string `gorm:"not null" form:"content" query:"content"  json:"content"`
	// 客户端名称
	AppId string `gorm:"not null" form:"app_id" query:"app_id" json:"app_id"`
	// 客户端版本号
	Version string `gorm:"not null" form:"version" query:"version" json:"version"`
	// 客户端 IP
	Ip string `gorm:"not null" form:"ip" query:"ip" json:"ip"`
	// 客户端用户信息
	User string `form:"user" query:"user" json:"user"`
	// 服务返回信息
	Serves string `form:"serves" query:"serves" json:"serves"`
	// 客户提交信息
	Client    string `form:"client" query:"client" json:"client"`
	CreatedAt time.Time
}

type AuthKey struct {
	Id  int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Key string `form:"key" json:"key"`
}
