package main

import "time"

type RecordLogs struct {
	Id        int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Type      string `gorm:"not null" form:"type" query:"type" json:"type"`
	Model     string `gorm:"not null" form:"model" query:"model" json:"model"`
	Content   string `gorm:"not null" form:"content" query:"content"  json:"content"`
	AppId     string `gorm:"not null" form:"app_id" query:"app_id" json:"app_id"`
	Ip        string `gorm:"not null" form:"ip" query:"ip" json:"ip"`
	User      string `form:"user" query:"user" json:"user"`
	Serves    string `form:"serves" query:"serves" json:"serves"`
	Client    string `form:"client" query:"client" json:"client"`
	CreatedAt time.Time
}

type AuthKey struct {
	Id  int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Key string `form:"key" json:"key"`
}
